package ctx

import (
	"context"
	"time"
)

type (
	operationIDKey string
	timeKey string
)

const OperationIDKey operationIDKey = "OperationID"
const TimeKey timeKey = "Time"

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

func GetTime(ctx context.Context) time.Time {
	value := ctx.Value(TimeKey)
	if value != nil {
		return value.(time.Time)
	}

	return time.Now().UTC()
}
