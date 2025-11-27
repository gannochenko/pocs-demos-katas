package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"worker/internal/github"
	"worker/internal/interfaces"
	"worker/internal/openai"
	"worker/internal/service/config"
	"worker/internal/service/monitoring"
	"worker/internal/slack"
	"worker/internal/temporal"
	"worker/internal/temporal/activities"
	"worker/internal/temporal/workflows"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"lib/logger"
	libMiddleware "lib/middleware"
	libTemporal "lib/temporal"
	"lib/util"
	"worker/internal/controller"
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

	// Register system endpoints
	healthHandler := controller.NewHealthHandler()
	e.GET("/health", healthHandler.Health)
	e.GET("/metrics", echo.WrapHandler(monitoringService.GetHandler()))

	// Register pprof endpoints
	// todo: remove for production
	// Open http://localhost:2024/debug/pprof/
	// To see the UI: go tool pprof -http=:8080 http://localhost:2024/debug/pprof/profile
	e.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))

	quitCh := make(chan struct{})

	temporalClient, err := libTemporal.GetTemporalClient(ctx, log, configService.Config.Temporal.ToClientOptions())
	if err != nil {
		return errors.Wrap(err, "could not get temporal client")
	}

	githubClient := github.NewClient(configService.GetConfig())
	githubClient.Connect(ctx)

	openaiClient := openai.NewClient(configService.GetConfig())

	slackClient := slack.NewClient(configService.GetConfig())

	startWorker, stopWorker, err := temporal.CreateWorker(
		ctx,
		temporalClient,
		log,
		configService.GetConfig(),
		[]interfaces.TemporalWorkflowGroup{
			workflows.NewReportWorkflowGroup(),
		},
		[]interfaces.TemporalActivityGroup{
			activities.NewReportActivityGroup(configService.GetConfig(), githubClient, openaiClient, slackClient),
		},
		quitCh,
	)

	return util.Run(ctx, quitCh, func(_ chan os.Signal) error {
		if err := startWorker(); err != nil {
			return errors.Wrap(err, "could not start temporal worker")
		}

		logger.Info(ctx, log, "Application started")

		return nil
	}, func() {
		cancel()
		stopWorker()

		err := e.Shutdown(ctx)
		if err != nil {
			logger.Error(ctx, log, errors.Wrap(err, "could not shutdown echo server").Error())
		}

		temporalClient.Close()
		monitoringService.Stop()

		logger.Info(ctx, log, "Application stopped")
	})
}

func main() {
	err := run(os.Stdout)
	if err != nil {
		logger.Error(context.TODO(), slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))

		os.Exit(1)
	}
}
