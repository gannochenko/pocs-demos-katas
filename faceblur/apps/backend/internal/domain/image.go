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
	IsFailed    bool
}

type ImageProcessingQueue struct {
	ID            uuid.UUID
	ImageID       uuid.UUID
	CreatedBy     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CompletedAt   *time.Time
	IsFailed      bool
	FailureReason *string
	OperationID   *string
}

type ListImagesRequest struct {
	PageNavigation PageNavigationRequest
}

type ListImagesResponse struct {
	Images         []Image
	PageNavigation PageNavigationResponse
}
