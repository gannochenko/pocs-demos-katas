package interfaces

import (
	"context"

	"backend/internal/domain"
	"backend/internal/util/logger"
)

type ImageService interface {
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
