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

	"google.golang.org/grpc"

	httpUtil "backend/internal/http"
	"backend/internal/service/config"
	"backend/internal/service/logger"
	loggerUtil "backend/internal/util/logger"
	"backend/internal/util/syserr"
)

func run(w io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// -----

	configService := config.NewConfigService()
	configuration, err := configService.GetConfig()
	if err != nil {
		return err
	}

	loggerService := logger.NewService(w)

	// ---

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configuration.GRPCPort))
	if err != nil {
		return err
	}

	opts := grpc.ChainUnaryInterceptor(
		//s.auth.PopulateUser,
		//request.PopulateContext(),
	)
	grpcServer := grpc.NewServer(opts)

	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	mux, shutdownGrpcClient, err := httpUtil.GetMux(ctx, configuration)
	if err != nil {
		return err
	}
	defer func() {
		err = shutdownGrpcClient()
		if err != nil {
			loggerService.LogError(ctx, syserr.Wrap(err, "could not shutdown gRPC client"))
		}
	}()

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
