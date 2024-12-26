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

	var wg sync.WaitGroup
	//wg.Add(2)
	wg.Add(1)

	//go func() {
	//	defer wg.Done()
	//	loggerService.Info(ctx, "gRPC server starting")
	//	shutdownGRPCServer, shutdownError := network.StartGRPCServer(ctx, configuration, &network.Controllers{
	//		ImageServiceV1: v1.NewImageController(loggerService),
	//	})
	//	loggerService.Info(ctx, "gRPC server started")
	//	if shutdownError != nil {
	//		loggerService.LogError(ctx, syserr.Wrap(shutdownError, "could not start gRPC server"))
	//	}
	//	if shutdownGRPCServer != nil {
	//		shutdownGRPCServer()
	//	}
	//
	//	loggerService.Info(ctx, "gRPC server stopped")
	//}()

	//grpcConnection, closeGPRCConnection, err := network.ConnectToGRPCServer(configuration)
	//defer func() {
	//	if closeGPRCConnection != nil {
	//		localErr := closeGPRCConnection()
	//		if localErr != nil {
	//			loggerService.LogError(ctx, syserr.Wrap(localErr, "could not close gRPC connection"))
	//		}
	//	}
	//}()
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Printf("%v", grpcConnection)
	//
	//mux, err := network.GetMux(ctx, grpcConnection)
	//if err != nil {
	//	return syserr.Wrap(err, "could not create mux")
	//}

	//go func() {
	//	defer wg.Done()
	//	loggerService.Info(ctx, "HTTP server starting")
	//	shutdownHTTPServer, localErr := network.StartHTTPServer(ctx, configuration, mux)
	//	loggerService.Info(ctx, "HTTP server started")
	//	if localErr != nil {
	//		loggerService.LogError(ctx, syserr.Wrap(localErr, "could not start HTTP server"))
	//	}
	//	if shutdownHTTPServer != nil {
	//		shutdownErr := shutdownHTTPServer()
	//		if shutdownErr != nil {
	//			loggerService.LogError(ctx, syserr.Wrap(shutdownErr, "could not shutdown HTTP server"))
	//		}
	//	}
	//
	//	loggerService.Info(ctx, "HTTP server stopped")
	//}()

	loggerService.Info(ctx, fmt.Sprintf("service started, http port %d", configuration.HTTPPort))

	sig := make(chan os.Signal, 1)
	done := make(chan struct{}) // Channel to signal the main function to exit
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	wg.Wait()

	go func() {
		loggerService.Info(ctx, "waiting for signal")
		<-sig // Wait for the first signal
		loggerService.Info(ctx, "service shutting down")

		cancel() // Cancel the context to signal all goroutines to stop
		loggerService.Info(ctx, "cancelled")
		wg.Wait() // Wait for all goroutines to complete
		loggerService.Info(ctx, "waited")
		close(done) // Signal the main function to exit
		loggerService.Info(ctx, "closed")
	}()

	<-done // Wait for the done channel to close before exiting

	//cancel()

	//fmt.Printf("wait 1\n")
	//wg.Wait()
	//fmt.Printf("waited 2\n")

	return nil
}

//err := run(os.Stdout)
//if err != nil {
//	loggerUtil.Error(nil, slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))
//
//	os.Exit(1)
//}

func main() {
	err := run1()
	if err != nil {
		loggerUtil.Error(nil, slog.New(slog.NewJSONHandler(os.Stdout, nil)), fmt.Sprintf("could not start the application: %s", err.Error()))

		os.Exit(1)
	}
}

func run1() error {
	var wg sync.WaitGroup
	wg.Add(1)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		defer wg.Done() // Mark WaitGroup as done when signal is received
		fmt.Println("WAITING")
		s := <-sig
		fmt.Printf("Signal received: %s\n", s.String())
	}()

	wg.Wait()

	fmt.Printf("DONE\n")

	return nil
}
