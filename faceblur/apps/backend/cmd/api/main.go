package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	v1 "backend/internal/controller/image/v1"
	"backend/internal/network"
	"backend/internal/service/config"
	"backend/internal/service/logger"
	loggerUtil "backend/internal/util/logger"
	"backend/internal/util/syserr"
)

func run(w io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configService := config.NewConfigService()
	configuration, err := configService.GetConfig()
	if err != nil {
		return err
	}

	loggerService := logger.NewService(w)

	go func() {
		shutdownGRPCServer, err := network.StartGRPCServer(configuration, &network.Controllers{
			ImageServiceV1: &v1.ImageController{},
		})
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not start gRPC server"))
		}
		shutdownGRPCServer()
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
		return err
	}

	go func() {
		shutdownHTTPServer, err := network.StartHTTPServer(ctx, configuration, mux)
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not start HTTP server"))
		}
		err = shutdownHTTPServer()
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not shutdown HTTP server"))
		}
	}()

	loggerService.Info(ctx, fmt.Sprintf("service started, http port %d", configuration.HTTPPort))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	loggerService.Info(ctx, "service shutting down")

	cancel()

	return nil
}

func main() {
	err := run(os.Stdout)
	if err != nil {
		loggerUtil.Error(nil, slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))

		os.Exit(1)
	}
}
