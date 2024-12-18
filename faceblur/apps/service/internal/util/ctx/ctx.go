package ctx

import "context"

type (
	timestampIDKey string
	operationIDKey string
	userEmailKey   string
)

const TimestampIDKey timestampIDKey = "Timestamp"
const OperationIDKey operationIDKey = "OperationID"
const UserEmailKey userEmailKey = "UserEmail"

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

func WithUserEmail(ctx context.Context, userEmail string) context.Context {
	return context.WithValue(ctx, UserEmailKey, userEmail)
}

func GetUserEmail(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	value := ctx.Value(UserEmailKey)
	if value == nil {
		return ""
	}

	return value.(string)
}
