package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	var opts []grpc.DialOption
	err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HTTPPort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			address := l.Addr().String()
			fmt.Println("Listening at " + address)
			return ctx
		},
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	cancel()
	_ = server.Shutdown(ctx)
}
