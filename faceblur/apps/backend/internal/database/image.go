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
	IsProcessed bool      `gorm:"type:bool"`
}

func (p *Image) ToDomain() (*domain.Image, error) {
	result := &domain.Image{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ListParameters struct {
}

type CountParameters struct {
}
