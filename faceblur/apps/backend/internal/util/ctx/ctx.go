package ctx

import (
	"context"

	"backend/internal/domain"
)

type (
	operationIDKey string
	userKey        string
)

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

func GetUser(ctx context.Context) *domain.User {
	if ctx == nil {
		return nil
	}

	value := ctx.Value(UserKey)
	if value == nil {
		return nil
	}

	user := value.(domain.User)

	return &user
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
