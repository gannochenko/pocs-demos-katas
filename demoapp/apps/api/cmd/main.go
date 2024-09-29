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
	"api/internal/controller/book"
	bookRepository "api/internal/repository/book"
	"api/internal/service"
	bookService "api/internal/service/book"
	"api/internal/util"
	"api/internal/util/db"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session, err := db.Connect()
	if err != nil {
		panic(err)
	}

	booksRepo := &bookRepository.Repository{
		Session: session,
	}
	bookSvc := &bookService.Service{
		BookRepository: booksRepo,
	}
	bookController := book.Controller{
		BookService: bookSvc,
	}

	PetAPIService := service.NewPetAPIService()
	PetAPIController := controller.NewPetAPIController(PetAPIService)

	StoreAPIService := service.NewStoreAPIService()
	StoreAPIController := controller.NewStoreAPIController(StoreAPIService)

	router := mux.NewRouter()

	router.HandleFunc("/books", bookController.GetBooks)

	util.PopulateRouter(router, PetAPIController, StoreAPIController)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
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
