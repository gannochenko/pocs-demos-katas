package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"api/internal/controller"
	"api/internal/factory"
	"api/internal/service"
	"api/internal/util"
	"api/internal/util/db"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configService := service.NewConfigService()
	config, err := configService.GetConfig()
	if err != nil {
		panic(err)
	}

	session, err := db.Connect(config.Postgres.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	serviceFactory := factory.MakeServiceFactory(session)

	PetAPIController := controller.NewPetAPIController(serviceFactory.GetPetService())
	StoreAPIController := controller.NewStoreAPIController(serviceFactory.GetStoreService())

	router := mux.NewRouter()
	util.PopulateRouter(router, PetAPIController, StoreAPIController)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTPPort),
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			address := l.Addr().String()
			fmt.Println("Listening at " + address)
			ctx = context.WithValue(ctx, keyServerAddress, address)
			return ctx
		},
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	cancel()
	_ = server.Shutdown(ctx)
}
