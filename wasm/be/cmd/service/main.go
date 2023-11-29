package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
	mux := http.NewServeMux()
	//mux.Handle("/books", updaterMiddleware(http.HandlerFunc(bookController.GetBooks)))

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
