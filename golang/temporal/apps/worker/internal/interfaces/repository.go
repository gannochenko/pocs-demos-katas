package interfaces

import (
	"context"
	"worker/internal/database"

	"gorm.io/gorm"
)

type WebhookDeduplicator interface {
	IsEventProcessed(ctx context.Context, tx *gorm.DB, eventID string) (bool, error)
	SetEventProcessed(ctx context.Context, tx *gorm.DB, event *database.WebhookEvent) error
}
