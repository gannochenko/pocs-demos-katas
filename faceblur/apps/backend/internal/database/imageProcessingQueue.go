package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"backend/internal/domain"
)

func NewImageProcessingQueueFromDomain(p *domain.ImageProcessingQueue) (*ImageProcessingQueue, error) {
	result := &ImageProcessingQueue{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ImageProcessingQueue struct {
	gorm.Model
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	ImageID       uuid.UUID  `gorm:"type:uuid"`
	CreatedBy     uuid.UUID  `gorm:"type:uuid"`
	CreatedAt     time.Time  `gorm:"type:timestamptz"`
	UpdatedAt     time.Time  `gorm:"type:timestamptz"`
	CompletedAt   *time.Time `gorm:"type:timestamptz"`
	IsFailed      bool       `gorm:"type:bool"`
	FailureReason *string    `gorm:"type:varchar(255)"`
	OperationID   *string    `gorm:"type:varchar(32)"`
}

func (p *ImageProcessingQueue) ToDomain() (*domain.ImageProcessingQueue, error) {
	result := &domain.ImageProcessingQueue{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ImageProcessingQueueFilter struct {
	CreatedBy *uuid.UUID
	IsFailed *bool
	IsCompleted *bool
}

type ImageProcessingQueueListParameters struct {
	Filter     *ImageProcessingQueueFilter
	Pagination *Pagination
}

type ImageProcessingQueueCountParameters struct {
	Filter *ImageProcessingQueueFilter
}
