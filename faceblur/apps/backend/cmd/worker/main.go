package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"backend/factory/repository"
	"backend/factory/service"
	"backend/internal/network"
	"backend/internal/util/db"
	loggerUtil "backend/internal/util/logger"
	"backend/internal/util/syserr"
)

func run(w io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session, err := db.Connect(os.Getenv("POSTGRES_DB_DSN"))
	if err != nil {
		return syserr.Wrap(err, "could not connect to the database")
	}

	repositoryFactory := repository.NewRepositoryFactory(session)
	serviceFactory := service.NewServiceFactory(session, w, repositoryFactory)

	configuration, err := serviceFactory.GetConfigService().GetConfig()
	if err != nil {
		return syserr.Wrap(err, "could not get config")
	}

	loggerService := serviceFactory.GetLoggerService()
	imageProcessor := serviceFactory.GetImageProcessorService()
	eventBusService := serviceFactory.GetEventBusService()
	monitoringService := serviceFactory.GetMonitoringService()

	err = monitoringService.Start()
	if err != nil {
		return syserr.Wrap(err, "could not start monitoring service")
	}

	HTTPServer := network.NewWorkerHTTPServer(serviceFactory.GetConfigService(), monitoringService)

	loggerService.Info(ctx, "service started")

	var shutdownSequenceWg sync.WaitGroup
	shutdownSequenceWg.Add(3)

	go func() {
		shutdownSequenceWg.Done()
		localErr := eventBusService.Start(ctx)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start event bus"))
		}
	}()

	go func() {
		shutdownSequenceWg.Done()
		localErr := imageProcessor.Start(ctx)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start image processor"))
		}
	}()

	go func() {
		shutdownSequenceWg.Done()
		localErr := HTTPServer.Start(ctx, configuration)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start HTTP server"))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-sig
	loggerService.Info(ctx, fmt.Sprintf("signal received: %s, starting shutdown sequence", s.String()))

	cancel()
	monitoringService.Stop()

	err = imageProcessor.Stop()
	if err != nil {
		loggerService.LogError(ctx, syserr.Wrap(err, "could not stop image processor"))
	}
	err = eventBusService.Stop()
	if err != nil {
		loggerService.LogError(ctx, err)
	}

	shutdownSequenceWg.Wait()
	loggerService.Info(ctx, "shutdown complete")

	return nil
}

func main() {
	err := run(os.Stdout)
	if err != nil {
		loggerUtil.Error(nil, slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))

		os.Exit(1)
	}
}
