package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"api/internal/domain"
)

type Pet struct {
	gorm.Model
	ID        uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string           `db:"name"`
	Status    domain.PetStatus `db:"status"`
	PhotoUrls []string         `gorm:"type:TEXT[]" db:"photo_urls"`
}

//func (p *Pet) ToDomain() *domain.Pet {
//	return &domain.Pet{
//		ID:        p.ID.String(),
//		Title:     p.Title,
//		Author:    p.Author,
//		IssueYear: p.IssueYear,
//	}
//}
