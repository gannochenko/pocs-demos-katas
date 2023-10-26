package monitoring

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otMetric "go.opentelemetry.io/otel/metric"
)

const (
	MeterName              = "app"
	LatencyMSHistogramName = "latency_ms"
)

var (
	meter otMetric.Meter
)

func RecordSpan(ctx context.Context, metricName string, metricValue string) func() {
	startTime := time.Now()

	return func() {
		duration := time.Since(startTime)

		hist, err := getMeter().Int64Histogram(LatencyMSHistogramName)
		if err != nil {
			log.Printf("could not get histogram: %s\n", err)
			return
		}
		hist.Record(ctx, duration.Milliseconds(), otMetric.WithAttributeSet(attribute.NewSet(attribute.String(metricName, metricValue))))
	}
}

func getMeter() otMetric.Meter {
	if meter == nil {
		meter = otel.GetMeterProvider().Meter(MeterName)
	}

	return meter
}
