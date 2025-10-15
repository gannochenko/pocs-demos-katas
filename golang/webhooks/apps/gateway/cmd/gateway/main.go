package main

import (
	"context"
	"fmt"
	"gateway/internal/controller"
	webhooksV1Handlers "gateway/internal/controller/v1/webhooks"
	"gateway/internal/database"
	"gateway/internal/factory/repository"
	"gateway/internal/factory/service"
	webhooksV1 "gateway/internal/http/v1"
	"gateway/internal/middleware"
	"gateway/internal/service/config"
	"gateway/internal/service/monitoring"
	"io"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"lib/logger"
	libMiddleware "lib/middleware"
	"lib/util"
)

func run(w io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := slog.New(slog.NewJSONHandler(w, nil))
	configService := config.NewConfigService()

	if err := configService.LoadConfig(); err != nil {
		return errors.Wrap(err, "could not load config")
	}

	db := database.NewDatabase(&configService.Config.Database, log)
	closeDb, err := db.Connect();
	if err != nil {
		return errors.Wrap(err, "could not connect to database")
	}
	defer closeDb()

	repositoryFactory := repository.New(db.DB)
	serviceFactory := service.NewFactory(db.DB, repositoryFactory)

	e := echo.New()

	monitoringService := monitoring.NewService(configService)
	if err := monitoringService.Start(); err != nil {
		return errors.Wrap(err, "could not start monitoring service")
	}

	e.HTTPErrorHandler = libMiddleware.ErrorHandler(log)
	e.Use(libMiddleware.LoggerMiddleware(log))
	e.Use(middleware.ObservabilityMiddleware(monitoringService))
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"}, // todo: restrict to only the frontend domain
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

	// Register system endpoints
	healthHandler := controller.NewHealthHandler()
	e.GET("/health", healthHandler.Health)
	e.GET("/metrics", echo.WrapHandler(monitoringService.GetHandler()))

	// Register pprof endpoints
	// todo: remove for production
	// Open http://localhost:2024/debug/pprof/
	// To see the UI: go tool pprof -http=:8080 http://localhost:2024/debug/pprof/profile
	e.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))

	webhooksHandler := webhooksV1Handlers.NewWebhooksHandler(serviceFactory.GetWebhookService())
	webhooksV1.RegisterHandlers(e, webhooksHandler)

	return util.Run(ctx, func() error {
		go func() {
			if err := e.Start(configService.Config.HTTP.Addr); err != nil {
				logger.Error(ctx, log, err.Error())
			}
		}()

		return nil
	}, func() error {
		cancel()
		err := e.Shutdown(ctx)
		if err != nil {
			logger.Error(ctx, log, errors.Wrap(err, "could not shutdown echo server").Error())
		}

		monitoringService.Stop()

		if err := closeDb(); err != nil {
			logger.Error(ctx, log, errors.Wrap(err, "could not close database connection").Error())
		}

		return nil
	})
}

func main() {
	err := run(os.Stdout)
	if err != nil {
		logger.Error(context.TODO(), slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))

		os.Exit(1)
	}
}
