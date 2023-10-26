package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"

	"goapp/internal/monitoring"
)

const (
	metricLabelHandler = "Handler"
	keyServerAddress   = "serverAddress"
)

func GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	defer monitoring.RecordSpan(request.Context(), metricLabelHandler, "GetBooks")()

	responseWriter.WriteHeader(200)
	_, _ = responseWriter.Write([]byte("returning books - ok"))
}

func GetAuthors(responseWriter http.ResponseWriter, request *http.Request) {
	defer monitoring.RecordSpan(request.Context(), metricLabelHandler, "GetAuthors")()

	responseWriter.WriteHeader(200)
	_, _ = responseWriter.Write([]byte("returning authors - ok"))
}

func GetChapters(responseWriter http.ResponseWriter, request *http.Request) {
	defer monitoring.RecordSpan(request.Context(), metricLabelHandler, "GetChapters")()

	responseWriter.WriteHeader(200)
	_, _ = responseWriter.Write([]byte("returning chapters - ok"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", GetBooks)
	mux.HandleFunc("/authors", GetAuthors)
	mux.HandleFunc("/chapters", GetChapters)

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
