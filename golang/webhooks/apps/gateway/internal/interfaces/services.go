package interfaces

import (
	"context"
	"gateway/internal/domain"
	"net/http"

	otelMetric "go.opentelemetry.io/otel/metric"
)

type ConfigService interface {
	LoadConfig() error
	GetConfig() *domain.Config
}

// MonitoringService defines the interface for monitoring and metrics operations
type MonitoringService interface {
	Start() error
	Stop()
	GetHandler() http.Handler
	AddInt64Counter(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string) error
	RecordInt64Gauge(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string, options ...otelMetric.Int64GaugeOption) error
	RecordInt64Histogram(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string, options ...otelMetric.Int64HistogramOption) error
}

type WebhookService interface {
	HandleWebhook(ctx context.Context, webhook *domain.WebhookEvent) error
}
