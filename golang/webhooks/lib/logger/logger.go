package logger

import (
	"context"
	"log/slog"

	pkgCtx "lib/ctx"
)

type Field struct {
	key   string
	value any
}

type LoggerFn func(context.Context, *slog.Logger, string, ...*Field)

func F(key string, value any) *Field {
	return &Field{
		key:   key,
		value: value,
	}
}

func Warning(ctx context.Context, logger *slog.Logger, message string, fields ...*Field) {
	logger.Warn(message, convertFields(addOperationID(ctx, fields))...)
}

func Error(ctx context.Context, logger *slog.Logger, message string, fields ...*Field) {
	logger.Error(message, convertFields(addOperationID(ctx, fields))...)
}

func Info(ctx context.Context, logger *slog.Logger, message string, fields ...*Field) {
	logger.Info(message, convertFields(addOperationID(ctx, fields))...)
}

func addOperationID(ctx context.Context, fields []*Field) []*Field {
	if ctx == nil {
		return fields
	}

	operationID := pkgCtx.GetOperationID(ctx)
	if operationID != "" {
		return append(fields, F("operation_id", pkgCtx.GetOperationID(ctx)))
	}

	return fields
}

func convertFields(fields []*Field) []any {
	result := make([]any, 2*len(fields))

	index := 0
	for _, field := range fields {
		result[index] = field.key
		result[index+1] = field.value

		index += 2
	}

	return result
}
