package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetTag struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PetID     uuid.UUID      `gorm:"type:uuid"`
	TagID     uuid.UUID      `gorm:"type:uuid"`
}
