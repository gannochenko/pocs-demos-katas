package database

import "gorm.io/gorm"

type WebhookEvent struct {
	gorm.Model
	EventID        string `gorm:"uniqueIndex;not null"`
	EventTimestamp string `gorm:"not null"`
	EventType      string `gorm:"index;not null"`
	Payload        string `gorm:"type:jsonb"`
}