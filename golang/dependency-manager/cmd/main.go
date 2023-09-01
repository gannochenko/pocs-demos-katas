package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"dependencymanager/internal/controller/book"
	"dependencymanager/internal/manager/repository"
	"dependencymanager/internal/manager/service"
	"dependencymanager/internal/util/db"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	session, err := db.Connect()
	if err != nil {
		panic(err)
	}

	svcManager := service.New(session, repository.New(session))

	bookController := book.Controller{
		BookService: svcManager.GetBookService(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/books", bookController.GetBooks)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mux,
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

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
