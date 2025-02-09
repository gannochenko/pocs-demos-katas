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
	v1 "backend/internal/controller/image/v1"
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

	gRPCSever := network.NewGRPCServer(
		&network.GRPCControllers{
			ImageServiceV1: v1.NewImageController(
				loggerService,
				serviceFactory.GetImageService(),
				serviceFactory.GetStorageService(),
				serviceFactory.GetConfigService(),
			),
		},
		loggerService,
		serviceFactory.GetAuthService(),
		serviceFactory.GetUserService(),
	)
	websocketServer := network.NewWebsocketServer(
		serviceFactory.GetConfigService(),
		serviceFactory.GetAuthService(),
		serviceFactory.GetLoggerService(),
		serviceFactory.GetUserService(),
		serviceFactory.GetEventBusService(),
	)
	HTTPServer := network.NewHTTPServer(serviceFactory.GetConfigService(), websocketServer)

	eventBusService := serviceFactory.GetEventBusService()

	var shutdownSequenceWg sync.WaitGroup
	shutdownSequenceWg.Add(4)

	go func() {
		shutdownSequenceWg.Done()
		localErr := gRPCSever.Start(ctx, configuration)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start gRPC server"))
		}
	}()

	go func() {
		shutdownSequenceWg.Done()
		localErr := websocketServer.Start(ctx)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start websocket server"))
		}
	}()

	go func() {
		shutdownSequenceWg.Done()
		localErr := HTTPServer.Start(ctx, configuration)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start gRPC server"))
		}
	}()

	go func() {
		shutdownSequenceWg.Done()
		localErr := eventBusService.Start(ctx)
		if localErr != nil {
			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start event bus"))
		}
	}()

	// add more background tasks here if needed

	loggerService.Info(ctx, fmt.Sprintf("service started, HTTP port %d, gRPC port %d", configuration.Backend.HTTP.Port, configuration.Backend.GRPCPort))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-sig
	loggerService.Info(ctx, fmt.Sprintf("signal received: %s, starting shutdown sequence", s.String()))

	cancel()
	gRPCSever.Stop()
	err = websocketServer.Stop()
	if err != nil {
		loggerService.LogError(ctx, err)
	}
	err = HTTPServer.Stop(ctx)
	if err != nil {
		loggerService.LogError(ctx, err)
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
