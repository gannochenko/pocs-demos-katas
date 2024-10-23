package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `db:"username"`
}
