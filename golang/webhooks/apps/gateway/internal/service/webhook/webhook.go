package webhook

import (
	"context"
	"gateway/internal/domain"
	"gateway/internal/interfaces"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
	webhookDeduplicator interfaces.WebhookDeduplicator
}

func NewService(db *gorm.DB, webhookDeduplicator interfaces.WebhookDeduplicator) *Service {
	return &Service{db: db, webhookDeduplicator: webhookDeduplicator}
}

func (s *Service) HandleWebhook(ctx context.Context, webhook *domain.WebhookEvent) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	isProcessed, err := s.webhookDeduplicator.IsEventProcessed(ctx, tx, webhook.EventID)
	if err != nil {
		return errors.Wrap(err, "could not check if event is processed")
	}

	if isProcessed {
		// already processed, skip
		return nil
	}

	// todo: enqueue the event
	
	err = s.webhookDeduplicator.SetEventProcessed(ctx, tx, webhook.EventID)
	if err != nil {
		return errors.Wrap(err, "could not set event as processed")
	}

	tx.Commit()

	return nil
}
