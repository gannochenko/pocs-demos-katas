package main

import (
	"context"
	"net"
	"net/http"
	"sync"

	"levelsgorm/internal/controller/book"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	bookController := book.Controller{}

	mux := http.NewServeMux()
	mux.HandleFunc("/books", bookController.GetBooks)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
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
