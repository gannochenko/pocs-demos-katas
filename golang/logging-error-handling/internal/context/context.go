package context

import (
	"context"

	"github.com/google/uuid"
)

type (
	operationIDKey string
)

const OperationIDKey operationIDKey = "OperationID"

func WithOperationID(ctx context.Context, operationID string) context.Context {
	if operationID != "" {
		// make sure operationID is a valid UUID
		_, err := uuid.Parse(operationID)
		if err != nil {
			operationID = ""
		}
	}

	if operationID == "" {
		operationID = uuid.New().String()
	}

	return context.WithValue(ctx, OperationIDKey, operationID)
}

func GetOperationID(ctx context.Context) string {
	value := ctx.Value(OperationIDKey)
	if value == nil {
		return ""
	}

	return value.(string)
}
