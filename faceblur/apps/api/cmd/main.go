package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpUtil "api/internal/http"
)

const (
	keyServerAddress = "serverAddress"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux, err := httpUtil.GetMux(ctx)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			address := l.Addr().String()
			fmt.Println("Listening at " + address)
			return context.WithValue(ctx, keyServerAddress, address)
		},
	}

	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	cancel()
	return server.Shutdown(ctx)
}

func main() {
	err := run()
	if err != nil {
		// log here!

		os.Exit(1)
	}
}
