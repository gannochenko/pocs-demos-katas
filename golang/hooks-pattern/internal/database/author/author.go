package author

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `db:"name"`
	HasBooks bool      `db:"has_books"`
}
