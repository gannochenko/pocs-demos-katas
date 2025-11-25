package repository

import (
	"worker/internal/interfaces"
	"worker/internal/repository/webhook_deduplicator"

	"gorm.io/gorm"
)

type Factory struct {
	session             *gorm.DB
	webhookDeduplicator interfaces.WebhookDeduplicator
}

func New(session *gorm.DB) *Factory {
	return &Factory{session: session}
}

func (f *Factory) GetWebhookDeduplicator() interfaces.WebhookDeduplicator {
	if f.webhookDeduplicator == nil {
		f.webhookDeduplicator = webhook_deduplicator.New(f.session)
	}
	return f.webhookDeduplicator
}
