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

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		loggerService.Info(ctx, "gRPC server starting")
		shutdownGRPCServer, err := network.StartGRPCServer(configuration, &network.Controllers{
			ImageServiceV1: v1.NewImageController(loggerService),
		})
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not start gRPC server"))
		} else {
			shutdownGRPCServer()
		}

		loggerService.Info(ctx, "gRPC server stopped")
		wg.Done()
	}()

	grpcConnection, closeGPRCConnection, err := network.ConnectToGRPCServer(configuration)
	if err != nil {
		return err
	}
	defer func() {
		err = closeGPRCConnection()
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not close gRPC connection"))
		}
	}()

	mux, err := network.GetMux(ctx, grpcConnection)
	if err != nil {
		return syserr.Wrap(err, "could not create mux")
	}

	go func() {
		loggerService.Info(ctx, "HTTP server starting")
		shutdownHTTPServer, err := network.StartHTTPServer(ctx, configuration, mux)
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not start HTTP server"))
		} else {
			err = shutdownHTTPServer()
			if err != nil {
				loggerService.LogError(ctx, syserr.Wrap(err, "could not shutdown HTTP server"))
			}
		}

		loggerService.Info(ctx, "HTTP server stopped")
		wg.Done()
	}()

	loggerService.Info(ctx, fmt.Sprintf("service started, http port %d", configuration.HTTPPort))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	loggerService.Info(ctx, "service shutting down")

	cancel()
	wg.Wait()

	return nil
}

func main() {
	err := run(os.Stdout)
	if err != nil {
		loggerUtil.Error(nil, slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))

		os.Exit(1)
	}
}
