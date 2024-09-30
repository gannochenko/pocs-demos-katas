package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"api/internal/domain"
)

type Order struct {
	gorm.Model
	ID       uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4()"`
	PetID    uuid.UUID          `db:"petId"`
	Quantity int32              `db:"quantity"`
	ShipDate time.Time          `db:"shipDate"`
	Status   domain.OrderStatus `db:"status"`
	Complete bool               `db:"complete"`
}
