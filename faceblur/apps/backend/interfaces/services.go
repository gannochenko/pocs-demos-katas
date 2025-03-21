package interfaces

import (
	"context"
	"image"
	"io"
	"net/http"
	"time"

	"backend/internal/domain"
	"backend/internal/util/logger"

	otelMetric "go.opentelemetry.io/otel/metric"
)

type ImageService interface {
	SubmitImageForProcessing(ctx context.Context, handle SessionHandle, url string, uploadedAt *time.Time) (*domain.Image, error)
	ListImages(ctx context.Context, handle SessionHandle, request *domain.ListImagesRequest) (*domain.ListImagesResponse, error)
}

type LoggerService interface {
	Warning(ctx context.Context, message string, fields ...*logger.Field)
	Error(ctx context.Context, message string, fields ...*logger.Field)
	Info(ctx context.Context, message string, fields ...*logger.Field)
	LogError(ctx context.Context, err error, fields ...*logger.Field)
}

type ConfigService interface {
	GetConfig() (*domain.Config, error)
}

type StorageService interface {
	GetWriter(ctx context.Context, bucketName string, objectPath string) (io.WriteCloser, error)
	PrepareSignedURL(ctx context.Context, bucketName string, objectPath string, ttl time.Duration, method string, contentType string) (url string, err error)
	GetPublicURL(bucketName string, objectName string) string
}

type AuthService interface {
	ValidateToken(ctx context.Context, token string) (string, int64, error)
	ExtractToken(ctx context.Context) (string, error)
}

type UserService interface {
	GetUserBySUP(ctx context.Context, sessionHandle SessionHandle, sup string) (*domain.User, error)
}

type EventBusService interface {
	Start(ctx context.Context) error
	Stop() error
	TriggerEvent(event *domain.EventBusEvent) error
	AddEventListener(eventType domain.EventBusEventType, cb func(event *domain.EventBusEvent)) error
	RemoveEventListener(eventType domain.EventBusEventType, cb func(event *domain.EventBusEvent)) error
}

type ImageProcessorService interface {
	Start(ctx context.Context) error
	Stop() error
}

type FaceDetectionService interface {
	Detect(ctx context.Context, image image.Image) ([]*domain.BoundingBox, error)
}

type MonitoringService interface {
	GetHandler() http.Handler
	Start() error
	Stop()
	AddInt64Counter(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string) error
	RecordInt64Gauge(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string, options ...otelMetric.Int64GaugeOption) error
	RecordInt64Histogram(ctx context.Context, meterName string, counterName string, value int64, labelName, labelValue string, options ...otelMetric.Int64HistogramOption) error
}
