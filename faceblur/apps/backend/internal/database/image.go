package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"backend/internal/domain"
)

func NewImageFromDomain(p *domain.Image) (*Image, error) {
	result := &Image{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type Image struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	URL         *string   `gorm:"type:varchar(255)"`
	OriginalURL string    `gorm:"type:varchar(255)"`
	CreatedBy   uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time `gorm:"type:timestamptz"`
	UpdatedAt   time.Time `gorm:"type:timestamptz"`
	UploadedAt  time.Time `gorm:"type:timestamptz"`
	IsProcessed bool      `gorm:"type:bool"`
	IsFailed    bool      `gorm:"type:bool"`
}

func (p *Image) ToDomain() (*domain.Image, error) {
	result := &domain.Image{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ImageUpdate struct {
	ID          uuid.UUID
	URL         *FieldValue[*string]
	OriginalURL *FieldValue[*string]
	IsProcessed *FieldValue[*bool]
	IsFailed *FieldValue[*bool]
}

type ImageFilter struct {
	CreatedBy *uuid.UUID
	ID        *uuid.UUID
}

type ImageListParameters struct {
	Filter     *ImageFilter
	Pagination *Pagination
}

type ImageCountParameters struct {
	Filter *ImageFilter
}
