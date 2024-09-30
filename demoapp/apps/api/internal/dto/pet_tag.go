package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetTag struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name,omitempty"`
}
