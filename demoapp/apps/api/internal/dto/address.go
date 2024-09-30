package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Street string    `db:"street"`
	City   string    `db:"city"`
	State  string    `db:"state"`
	Zip    string    `db:"zip"`
}
