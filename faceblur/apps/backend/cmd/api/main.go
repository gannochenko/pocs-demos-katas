package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	shutdownGRPCServer, err := network.StartGRPCServer(configuration)
	if err != nil {
		return err
	}
	defer shutdownGRPCServer()

	fmt.Printf("1")

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

	fmt.Printf("2")

	mux, err := network.GetMux(ctx, grpcConnection)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", configuration.HTTPPort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			address := l.Addr().String()
			fmt.Println("Listening at " + address)
			return ctx
		},
	}

	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	defer func() {
		err = server.Shutdown(ctx)
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not shutdown HTTP server"))
		}
	}()

	fmt.Printf("3")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	cancel()

	return nil
}

func main() {
	err := run(os.Stdout)
	if err != nil {
		loggerUtil.Error(nil, slog.New(slog.NewJSONHandler(os.Stdout, nil)), "could not start the application")

		os.Exit(1)
	}
}
