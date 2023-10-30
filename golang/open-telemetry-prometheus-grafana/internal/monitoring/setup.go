package monitoring

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	openTelemetryPrometheus "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func Setup(ctx context.Context, serviceName, serviceVersion string) (shutdown func(context.Context) error, prometheusRegistry *prometheus.Registry, err error) {
	var shutdownFuncs []func(context.Context) error

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

	res := resource.Environment()

	//// TODO: Setup trace provider
	//tracerProvider, err := newTraceProvider(res)
	//if err != nil {
	//	handleErr(err)
	//	return
	//}
	//shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	//otel.SetTracerProvider(tracerProvider)

	prometheusRegistry = createPrometheusRegistry()

	// Setup meter provider.
	meterProvider, err := createMeterProvider(prometheusRegistry, res) // res
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return
}

func SetupHTTP(mux *http.ServeMux, prometheusRegistry *prometheus.Registry) {
	Handler := func() http.Handler {
		return promhttp.InstrumentMetricHandler(
			prometheusRegistry, promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{}),
		)
	}

	mux.Handle("/metrics", Handler())
}

func createPrometheusRegistry() *prometheus.Registry {
	registry := prometheus.NewRegistry()

	registry.MustRegister(collectors.NewBuildInfoCollector())
	registry.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	return registry
}

func createMeterProvider(reg prometheus.Registerer, resource *resource.Resource) (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	exporter, err := openTelemetryPrometheus.New(
		openTelemetryPrometheus.WithRegisterer(reg),
		//openTelemetryPrometheus.WithAggregationSelector(histogramAggregationSelector)
	)
	if err != nil {
		return nil, fmt.Errorf("could not create prometheus exporter: %w", err)
	}

	provider := metric.NewMeterProvider(
		metric.WithResource(resource),

		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			metric.WithInterval(5*time.Second))),

		metric.WithReader(exporter),
	)
	return provider, nil
}
