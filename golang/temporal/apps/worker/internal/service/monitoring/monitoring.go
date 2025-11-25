package monitoring

import (
	"context"
	"net/http"
	"regexp"
	"worker/internal/interfaces"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	openTelemetryPrometheus "go.opentelemetry.io/otel/exporters/prometheus"
	otelMetric "go.opentelemetry.io/otel/metric"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type Service struct {
	configService interfaces.ConfigService
	meterProvider *metric.MeterProvider
	prometheusRegistry *prometheus.Registry
}

func NewService(configService interfaces.ConfigService) *Service {
	return &Service{
		configService: configService,
	}
}

func (s *Service) Start() error {
	config := s.configService.GetConfig()

	res, err := resource.Merge(resource.Environment(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(config.Service.Name),
			semconv.ServiceVersion(config.Service.Version),
		))
	if err != nil {
		return errors.Wrap(err, "could not create resource")
	}

	s.prometheusRegistry = s.createPrometheusRegistry()

	meterProvider, err := s.createMeterProvider(s.prometheusRegistry, res)
	if err != nil {
		return errors.Wrap(err, "could not create meter provider")
	}
	otel.SetMeterProvider(meterProvider)
	s.meterProvider = meterProvider

	return nil
}

func (s *Service) Stop() {
	if s.meterProvider != nil {
		s.meterProvider.Shutdown(context.Background())
	}
}

func (s *Service) GetHandler() http.Handler {
	return promhttp.InstrumentMetricHandler(
		s.prometheusRegistry, promhttp.HandlerFor(s.prometheusRegistry, promhttp.HandlerOpts{}),
	)
}

func (s *Service) AddInt64Counter(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string) error {
	counter, err := otel.GetMeterProvider().Meter(meterName).Int64Counter(counterName)
	if err != nil {
		return errors.Wrap(err, "could not create counter")
	}

	var opts []otelMetric.AddOption
	if labelName != "" {
		opts = append(opts, otelMetric.WithAttributeSet(attribute.NewSet(attribute.String(labelName, labelValue))))
	}

	counter.Add(ctx, value, opts...)

	return nil
}

func (s *Service) RecordInt64Gauge(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string, options ...otelMetric.Int64GaugeOption) error {
	counter, err := otel.GetMeterProvider().Meter(meterName).Int64Gauge(counterName, options...)
	if err != nil {
		return errors.Wrap(err, "could not create gauge")
	}

	var opts []otelMetric.RecordOption
	if labelName != "" {
		opts = append(opts, otelMetric.WithAttributeSet(attribute.NewSet(attribute.String(labelName, labelValue))))
	}

	counter.Record(ctx, value, opts...)

	return nil
}

func (s *Service) RecordInt64Histogram(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string, options ...otelMetric.Int64HistogramOption) error {
	counter, err := otel.GetMeterProvider().Meter(meterName).Int64Histogram(counterName, options...)
	if err != nil {
		return errors.Wrap(err, "could not create histogram")
	}

	var opts []otelMetric.RecordOption
	if labelName != "" {
		opts = append(opts, otelMetric.WithAttributeSet(attribute.NewSet(attribute.String(labelName, labelValue))))
	}

	counter.Record(ctx, value, opts...)

	return nil
}

func (s *Service) createPrometheusRegistry() *prometheus.Registry {
	registry := prometheus.NewRegistry()

	registry.MustRegister(collectors.NewBuildInfoCollector())
	registry.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	return registry
}

func (s *Service) createMeterProvider(reg prometheus.Registerer, resource *resource.Resource) (*metric.MeterProvider, error) {
	exporter, err := openTelemetryPrometheus.New(
		openTelemetryPrometheus.WithRegisterer(reg),
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not create prometheus exporter")
	}

	provider := metric.NewMeterProvider(
		metric.WithResource(resource),
		metric.WithReader(exporter),
	)
	return provider, nil
}
