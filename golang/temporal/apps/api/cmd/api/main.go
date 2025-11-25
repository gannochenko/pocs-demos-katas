package main

import (
	"api/internal/client/temporal"
	"api/internal/controller"
	"api/internal/factory"
	"api/internal/middleware"
	"api/internal/service/config"
	"api/internal/service/monitoring"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	workflowsV1Handlers "api/internal/controller/v1/workflows"
	workflowsV1 "api/internal/http/v1"
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

	temporalClient, err := temporal.GetTemporalClient(configService.Config)
	if err != nil {
		return errors.Wrap(err, "could not get temporal client")
	}

	depFactory := factory.NewFactory()
	depFactory.SetTemporalClient(temporalClient)

	// Register routes
	workflowsHandler := workflowsV1Handlers.NewWorkflowsHandler(depFactory.GetTemporalService())
	workflowsV1.RegisterHandlers(e, workflowsHandler)

	return util.Run(ctx, nil, func(sigChan chan os.Signal) error {
		go func() {
			if err := e.Start(configService.Config.HTTP.Addr); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					// this is just a normal shutdown, so we can ignore it
					return
				}
				logger.Error(ctx, log, err.Error())
				sigChan <- nil // server couldn't start, exiting
			}
		}()

		logger.Info(ctx, log, "Application started")

		return nil
	}, func() error {
		cancel()
		err := e.Shutdown(ctx)
		if err != nil {
			logger.Error(ctx, log, errors.Wrap(err, "could not shutdown echo server").Error())
		}

		monitoringService.Stop()

		logger.Info(ctx, log, "Application stopped")

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
