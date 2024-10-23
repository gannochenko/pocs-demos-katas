package dto

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"api/internal/domain"
)

type Tag struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name,omitempty"`
}

func (t *Tag) ToDomain() (*domain.Tag, error) {
	result := &domain.Tag{}
	err := copier.Copy(result, t)
	if err != nil {
		return nil, err
	}

	return result, nil
}
