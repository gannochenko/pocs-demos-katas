package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"backend/internal/domain"
)

func NewUserFromDomain(p *domain.User) (*User, error) {
	result := &User{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"type:varchar(255)"`
	Sup       string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamptz"`
	UpdatedAt time.Time `gorm:"type:timestamptz"`
}

func (p *User) ToDomain() (*domain.User, error) {
	result := &domain.User{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}
