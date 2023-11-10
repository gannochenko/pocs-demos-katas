package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"hookspattern/internal/controller/book"
	authorRepository "hookspattern/internal/repository/author"
	bookRepository "hookspattern/internal/repository/book"
	bookService "hookspattern/internal/service/book"
	"hookspattern/internal/service/hooks"
	"hookspattern/internal/service/updater"
	"hookspattern/internal/util/db"
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
	authorRepo := &authorRepository.Repository{
		Session: session,
	}

	hooksSvc := &hooks.Service{}
	bookSvc := &bookService.Service{
		BookRepository: booksRepo,
		HooksService:   hooksSvc,
	}
	updaterSvc := updater.New(hooksSvc, authorRepo)

	bookController := book.Controller{
		BookService: bookSvc,
	}

	updaterSvc.Init()

	updaterMiddleware := updaterSvc.GetHTTPMiddleware()

	mux := http.NewServeMux()
	mux.Handle("/books", updaterMiddleware(http.HandlerFunc(bookController.GetBooks)))
	mux.Handle("/books/delete", updaterMiddleware(http.HandlerFunc(bookController.DeleteBook)))

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
