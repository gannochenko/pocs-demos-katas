package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"api/internal/domain"
	"api/internal/util/db"
)

type Pet struct {
	gorm.Model
	ID        uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string           `db:"name"`
	Status    domain.PetStatus `db:"status"`
	PhotoUrls []string         `gorm:"type:text[]"`
}

//func (p *Pet) ToDomain() *domain.Pet {
//	return &domain.Pet{
//		ID:        p.ID.String(),
//		Title:     p.Title,
//		Author:    p.Author,
//		IssueYear: p.IssueYear,
//	}
//}

type ListPetFilter struct {
	ID *string
}

type ListPetParameters struct {
	Filter     *ListPetFilter
	Pagination *db.Pagination
}
