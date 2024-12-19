package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpUtil "backend/internal/http"
	"backend/internal/service/config"
	"backend/internal/util"
	"backend/internal/util/syserr"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configService := config.NewConfigService()
	configuration, err := configService.GetConfig()
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
			util.LogError(ctx, syserr.Wrap(err, "could not shutdown gRPC client"))
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
			util.LogError(ctx, syserr.Wrap(err, "could not shutdown HTTP server"))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	cancel()

	return nil
}

func main() {
	err := run()
	if err != nil {
		util.LogError(nil, syserr.Wrap(err, "could not start the application"))

		os.Exit(1)
	}
}
