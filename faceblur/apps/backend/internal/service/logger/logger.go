package logger

import (
	"context"
	"io"
	"log/slog"

	"backend/internal/util/logger"
	"backend/internal/util/syserr"
)

type Service struct {
	logger *slog.Logger
}

var (
	errorToLogLevel = map[syserr.Code]func(ctx context.Context, logger *slog.Logger, message string, fields ...*logger.Field){
		syserr.InternalCode:       logger.Error,
		syserr.BadInputCode:       logger.Warning,
		syserr.NotFoundCode:       logger.Warning,
		syserr.NotImplementedCode: logger.Error,
	}
)

func (s *Service) Warning(ctx context.Context, message string, fields ...*logger.Field) {
	logger.Warning(ctx, s.logger, message, fields...)
}

func (s *Service) Error(ctx context.Context, message string, fields ...*logger.Field) {
	logger.Error(ctx, s.logger, message, fields...)
}

func (s *Service) Info(ctx context.Context, message string, fields ...*logger.Field) {
	logger.Info(ctx, s.logger, message, fields...)
}

func (s *Service) LogError(ctx context.Context, err error, fields ...*logger.Field) {
	code := syserr.GetCode(err)
	fn := errorToLogLevel[code]
	if fn == nil {
		fn = logger.Error
	}

	fields = append(fields, s.convertErrorFieldsToLoggerFields(syserr.GetFields(err))...)
	fields = append(fields, logger.F("stack", syserr.GetStackFormatted(err)))

	fn(ctx, s.logger, err.Error(), fields...)
}

func (s *Service) convertErrorFieldsToLoggerFields(fields []*syserr.Field) []*logger.Field {
	result := make([]*logger.Field, len(fields))

	for index, field := range fields {
		result[index] = logger.F(field.Key, field.Value)
	}

	return result
}

func NewLoggerService(logWriter io.Writer) *Service {
	return &Service{
		logger: slog.New(slog.NewJSONHandler(logWriter, nil)),
	}
}
