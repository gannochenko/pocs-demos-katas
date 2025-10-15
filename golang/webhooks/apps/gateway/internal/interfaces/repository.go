package interfaces

import (
	"context"

	"gorm.io/gorm"
)

type WebhookDeduplicator interface {
	IsEventProcessed(ctx context.Context, tx *gorm.DB, eventID string) (bool, error)
	SetEventProcessed(ctx context.Context, tx *gorm.DB, eventID string) error
}
