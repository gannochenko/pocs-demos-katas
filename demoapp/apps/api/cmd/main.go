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

	petV3 "api/internal/controller/pet/v3"
	storeV3 "api/internal/controller/store/v3"
	"api/internal/factory"
	"api/internal/service/config"
	"api/internal/util"
	"api/internal/util/db"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configService := config.NewConfigService()
	conf, err := configService.GetConfig()
	if err != nil {
		panic(err)
	}

	session, err := db.Connect(conf.Postgres.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	serviceFactory := factory.MakeServiceFactory(session)

	router := mux.NewRouter()
	util.PopulateRouter(
		router,
		petV3.NewPetAPIController(serviceFactory.GetPetService()),
		storeV3.NewStoreAPIController(serviceFactory.GetStoreService()),
	)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HTTPPort),
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
