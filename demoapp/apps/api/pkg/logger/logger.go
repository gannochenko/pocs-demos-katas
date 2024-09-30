package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

type Field struct {
	key   string
	value any
}

func F(key string, value any) *Field {
	return &Field{
		key:   key,
		value: value,
	}
}

func Warning(ctx context.Context, message string, fields ...*Field) {
	logger.Warn(message, convertFields(addOperationID(ctx, fields))...)
}

func Error(ctx context.Context, message string, fields ...*Field) {
	logger.Error(message, convertFields(addOperationID(ctx, fields))...)
}

func Info(ctx context.Context, message string, fields ...*Field) {
	logger.Info(message, convertFields(addOperationID(ctx, fields))...)
}

func addOperationID(ctx context.Context, fields []*Field) []*Field {
	return fields
	//return append(fields, F("operation_id", pkgContext.GetOperationID(ctx)))
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
