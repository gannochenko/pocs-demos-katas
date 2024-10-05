package ctx

import "context"

type (
	timestampIDKey string
	operationIDKey string
)

const TimestampIDKey timestampIDKey = "Timestamp"
const OperationIDKey operationIDKey = "OperationID"

func GetOperationID(ctx context.Context) string {
	value := ctx.Value(OperationIDKey)
	if value != nil {
		return value.(string)
	}

	return ""
}

func WithOperationID(ctx context.Context, operationID string) context.Context {
	if operationID == "" {
		return ctx
	}

	return context.WithValue(ctx, OperationIDKey, operationID)
}
