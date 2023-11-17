package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"goapp/internal/config"
	"goapp/internal/monitoring"
)

const (
	metricLabelHandler = "Handler"
	keyServerAddress   = "serverAddress"
)

func QueryBooks(ctx context.Context) {
	defer monitoring.RecordSpan(ctx, metricLabelHandler, "QueryBooks")()

	randomDelay()
}

func LoadBooks(ctx context.Context) {
	defer monitoring.RecordSpan(ctx, metricLabelHandler, "LoadBooks")()

	randomDelay()
	QueryBooks(ctx)
	randomDelay()
}

func GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	defer monitoring.RecordSpan(request.Context(), metricLabelHandler, "GetBooks")()

	randomDelay()
	LoadBooks(request.Context())
	randomDelay()

	responseWriter.WriteHeader(200)
	_, _ = responseWriter.Write([]byte("returning books - ok"))
}

func GetAuthors(responseWriter http.ResponseWriter, request *http.Request) {
	defer monitoring.RecordSpan(request.Context(), metricLabelHandler, "GetAuthors")()

	randomDelay()

	responseWriter.WriteHeader(200)
	_, _ = responseWriter.Write([]byte("returning authors - ok"))
}

func GetChapters(responseWriter http.ResponseWriter, request *http.Request) {
	defer monitoring.RecordSpan(request.Context(), metricLabelHandler, "GetChapters")()

	randomDelay()

	responseWriter.WriteHeader(200)
	_, _ = responseWriter.Write([]byte("returning chapters - ok"))
}

func randomDelay() {
	min := 1
	max := 2

	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(delay) * time.Second)
}

func main() {
	conf := &config.Config{
		Country:        os.Getenv("COUNTRY"),
		ServiceName:    os.Getenv("SERVICE_NAME"),
		ServiceVersion: os.Getenv("SERVICE_VERSION"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/books", GetBooks)
	mux.HandleFunc("/authors", GetAuthors)
	mux.HandleFunc("/chapters", GetChapters)

	ctx := context.Background()

	shutdownMonitoring, prometheusRegistry, err := monitoring.Setup(ctx, conf)
	if err != nil {
		os.Exit(1)
	}

	defer shutdownMonitoring(ctx)

	monitoring.SetupHTTP(mux, prometheusRegistry)

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
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
