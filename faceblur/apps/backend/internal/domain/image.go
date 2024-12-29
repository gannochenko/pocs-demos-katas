package domain

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID          uuid.UUID
	URL         *string
	OriginalURL string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsProcessed bool
}
