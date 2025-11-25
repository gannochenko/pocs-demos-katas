package ctx

import (
	"context"
)

type (
	operationIDKey string
)

const OperationIDKey operationIDKey = "OperationID"

func WithOperationID(ctx context.Context, operationID string) context.Context {
	if operationID == "" {
		return ctx
	}

	return context.WithValue(ctx, OperationIDKey, operationID)
}

func GetOperationID(ctx context.Context) string {
	value := ctx.Value(OperationIDKey)
	if value != nil {
		return value.(string)
	}

	return ""
}

func IsDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
	}

	return false
}

func IsTimeouted(ctx context.Context) bool {
	return ctx.Err() == context.DeadlineExceeded
}
