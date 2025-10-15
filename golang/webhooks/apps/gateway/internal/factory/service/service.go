package service

import (
	"gateway/internal/factory/repository"
	"gateway/internal/interfaces"
	"gateway/internal/service/webhook"

	"gorm.io/gorm"
)

type Factory struct {
	db *gorm.DB
	repositoryFactory *repository.Factory

	webhookService interfaces.WebhookService
}

func NewFactory(db *gorm.DB, repositoryFactory *repository.Factory) *Factory {
	return &Factory{db: db, repositoryFactory: repositoryFactory}
}

func (f *Factory) GetWebhookService() interfaces.WebhookService {
	if f.webhookService == nil {
		f.webhookService = webhook.NewService(f.db, f.repositoryFactory.GetWebhookDeduplicator())
	}
	return f.webhookService
}
