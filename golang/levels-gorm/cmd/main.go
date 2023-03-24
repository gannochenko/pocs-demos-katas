package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"levelsgorm/internal/controller/book"
	bookService "levelsgorm/internal/service/book"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	bookSvc := &bookService.Service{}

	bookController := book.Controller{
		BookService: bookSvc,
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

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
