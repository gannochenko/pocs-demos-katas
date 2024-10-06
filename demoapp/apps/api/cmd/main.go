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

	categoryV3 "api/internal/controller/category/v3"
	petV3 "api/internal/controller/pet/v3"
	tagV3 "api/internal/controller/tag/v3"
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

	serviceFactory := factory.MakeServiceFactory(session, configService)

	router := mux.NewRouter()
	util.PopulateRouter(
		router,
		conf,
		serviceFactory.GetAuthService(),
		petV3.NewPetAPIController(serviceFactory.GetPetService()),
		tagV3.NewTagAPIController(serviceFactory.GetTagService()),
		categoryV3.NewCategoryAPIController(serviceFactory.GetCategoryService()),
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
