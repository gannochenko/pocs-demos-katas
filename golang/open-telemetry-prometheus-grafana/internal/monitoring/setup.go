package monitoring

import (
	"context"
	"errors"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
)

func setupSDK(ctx context.Context, serviceName, serviceVersion string) (shutdown func(context.Context) error, promRegistry *prometheus.Registry, err error) {
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
