package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"loggingerrorhandling/internal/controller/book"
	httpInternal "loggingerrorhandling/internal/http"
	bookRepository "loggingerrorhandling/internal/repository/book"
	bookService "loggingerrorhandling/internal/service/book"
	"loggingerrorhandling/internal/util/db"
)

const (
	keyServerAddress = "serverAddress"
)

func main() {
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

	mux := http.NewServeMux()
	mux.Handle("/books", httpInternal.ResponseWriter(bookController.GetBooks))

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
