package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetTag struct {
	gorm.Model
	PetID uuid.UUID `gorm:"type:uuid"`
	TagID uuid.UUID `gorm:"type:uuid"`
}
