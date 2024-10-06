package dto

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"api/internal/domain"
)

type Category struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name,omitempty"`
}

func (c *Category) ToDomain() (*domain.Category, error) {
	result := &domain.Category{}
	err := copier.Copy(result, c)
	if err != nil {
		return nil, err
	}

	return result, nil
}
