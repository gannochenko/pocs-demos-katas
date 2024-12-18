package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpUtil "service/internal/http"
	"service/internal/util"
	"service/internal/util/syserr"
)

//const (
//	keyServerAddress = "serverAddress"
//)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux, shutdownGrpcClient, err := httpUtil.GetMux(ctx)
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
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			address := l.Addr().String()
			fmt.Println("Listening at " + address)
			//return context.WithValue(ctx, keyServerAddress, address)
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
		// log here!

		os.Exit(1)
	}
}
