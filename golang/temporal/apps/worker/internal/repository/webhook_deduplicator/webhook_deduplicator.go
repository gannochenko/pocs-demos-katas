package webhook_deduplicator

import (
	"context"

	"worker/internal/database"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repository struct {
	session *gorm.DB
}

func New(session *gorm.DB) *Repository {
	return &Repository{session: session}
}

func (r *Repository) IsEventProcessed(ctx context.Context, tx *gorm.DB, eventID string) (bool, error) {
	db := r.getRunner(tx)

	var event database.WebhookEvent
	err := db.WithContext(ctx).Where("event_id = ?", eventID).First(&event).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrap(err, "could not query webhook event")
	}

	return true, nil
}

func (r *Repository) SetEventProcessed(ctx context.Context, tx *gorm.DB, event *database.WebhookEvent) error {
	db := r.getRunner(tx)

	err := db.WithContext(ctx).Create(&event).Error
	if err != nil {
		return errors.Wrap(err, "could not create webhook event record")
	}

	return nil
}

func (r *Repository) getRunner(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.session
}
