package ctx

import (
	"context"

	"backend/internal/domain"
)

type (
	timestampIDKey string
	operationIDKey string
	userKey        string
)

const TimestampIDKey timestampIDKey = "Timestamp"
const OperationIDKey operationIDKey = "OperationID"
const UserKey userKey = "User"

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

func WithUser(ctx context.Context, user domain.User) context.Context {
	return context.WithValue(ctx, UserKey, user)
}

func GetUser(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	value := ctx.Value(UserKey)
	if value == nil {
		return ""
	}

	return value.(string)
}
