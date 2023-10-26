package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelProm "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	oM "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

/*
Combined examples:
1. https://github.com/open-telemetry/opentelemetry-go/blob/main/example/prometheus/main.go
2. https://opentelemetry.io/docs/instrumentation/go/getting-started/
*/

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context, serviceName, serviceVersion string) (shutdown func(context.Context) error, promRegistry *prometheus.Registry, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	//// Setup resource.
	//res, err := newResource(serviceName, serviceVersion)
	//if err != nil {
	//	handleErr(err)
	//	return
	//}

	res := resource.Environment()

	//// Setup trace provider.
	//tracerProvider, err := newTraceProvider(res)
	//if err != nil {
	//	handleErr(err)
	//	return
	//}
	//shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	//otel.SetTracerProvider(tracerProvider)

	promRegistry = createPrometheusRegistry()

	// Setup meter provider.
	meterProvider, err := newMeterProvider(promRegistry, res) // res
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return
}

// https://opentelemetry.io/docs/instrumentation/js/resources/
//func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
//	return resource.Merge(resource.Default(),
//		resource.NewWithAttributes(semconv.SchemaURL,
//			semconv.ServiceName(serviceName),
//			semconv.ServiceVersion(serviceVersion),
//		))
//}

//func newTraceProvider(res *resource.Resource) (*trace.TracerProvider, error) {
//	traceExporter, err := stdouttrace.New(
//		stdouttrace.WithPrettyPrint())
//	if err != nil {
//		return nil, err
//	}
//
//	traceProvider := trace.NewTracerProvider(
//		trace.WithBatcher(traceExporter,
//			// Default is 5s. Set to 1s for demonstrative purposes.
//			trace.WithBatchTimeout(time.Second)),
//		trace.WithResource(res),
//	)
//	return traceProvider, nil
//}

func createPrometheusRegistry() *prometheus.Registry {
	//prometheusRegisterer := prometheusClient.NewRegistry()
	reg := prometheus.NewRegistry()

	//reg.MustRegister(collectors.NewBuildInfoCollector())
	//reg.MustRegister(collectors.NewGoCollector(
	//	collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	//))

	return reg
}

func newMeterProvider(reg prometheus.Registerer, res *resource.Resource) (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	exporter, err := otelProm.New(
		otelProm.WithRegisterer(reg),
		//otelProm.WithAggregationSelector(histogramAggregationSelector)
	)
	if err != nil {
		return nil, fmt.Errorf("could not create prometheus exporter: %w", err)
	}

	provider := metric.NewMeterProvider(
		// resource это просто информация о сервисе
		metric.WithResource(res),

		// эта штука срет в консоль:
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(3*time.Second))),

		metric.WithReader(exporter),
	)
	return provider, nil
}

func run() (err error) {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	serviceName := "dice"
	serviceVersion := "0.1.0"
	otelShutdown, promRegistry, err := setupOTelSDK(ctx, serviceName, serviceVersion)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":1080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(promRegistry),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler(prometheusReg *prometheus.Registry) http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	Handler := func() http.Handler {
		return promhttp.InstrumentMetricHandler(
			prometheusReg, promhttp.HandlerFor(prometheusReg, promhttp.HandlerOpts{}),
		)
	}

	mux.Handle("/metrics", Handler())
	handleFunc("/rolldice", rolldice)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}

var (
	//tracer  = otel.Tracer("rolldice")
	meter   = otel.Meter("rolldice")
	rollCnt oM.Int64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter("dice.rolls",
		oM.WithDescription("The number of rolls by roll value"),
		oM.WithUnit("{roll}"))
	if err != nil {
		panic(err)
	}
}

func rolldice(w http.ResponseWriter, r *http.Request) {
	//ctx, span := tracer.Start(r.Context(), "roll")
	//defer span.End()
	ctx := context.Background()

	roll := 1 + rand.Intn(6)

	// Add the custom attribute to the span and counter.
	rollValueAttr := attribute.Int("roll.value", roll)
	//span.SetAttributes(rollValueAttr)
	rollCnt.Add(ctx, 1, oM.WithAttributes(rollValueAttr))

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
