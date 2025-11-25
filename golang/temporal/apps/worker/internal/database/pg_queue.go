package database

import (
	"time"

	"gorm.io/gorm"
)

type PGQueueStatus string

const (
	PGQueueStatusPending PGQueueStatus = "pending"
	PGQueueStatusProcessed PGQueueStatus = "processed"
	PGQueueStatusFailed PGQueueStatus = "failed"
)

type PGQueue struct {
	gorm.Model
	Status        PGQueueStatus `gorm:"type:pg_queue_status;not null"`
	Payload       string `gorm:"type:jsonb"`
	CreatedAt     time.Time  `gorm:"type:timestamptz"`
}