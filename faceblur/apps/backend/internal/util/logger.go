package util

import (
	"context"

	"backend/internal/util/logger"
	"backend/internal/util/syserr"
)

var (
	errorToLogLevel = map[syserr.Code]func(ctx context.Context, message string, fields ...*logger.Field){
		syserr.InternalCode:       logger.Error,
		syserr.BadInputCode:       logger.Warning,
		syserr.NotFoundCode:       logger.Warning,
		syserr.NotImplementedCode: logger.Error,
	}
)

func LogError(ctx context.Context, err error, fields ...*logger.Field) {
	code := syserr.GetCode(err)
	fn := errorToLogLevel[code]
	if fn == nil {
		fn = logger.Error
	}

	fields = append(fields, convertErrorFieldsToLoggerFields(syserr.GetFields(err))...)
	fields = append(fields, logger.F("stack", syserr.GetStackFormatted(err)))

	fn(ctx, err.Error(), fields...)
}

func convertErrorFieldsToLoggerFields(fields []*syserr.Field) []*logger.Field {
	result := make([]*logger.Field, len(fields))

	for index, field := range fields {
		result[index] = logger.F(field.Key, field.Value)
	}

	return result
}
