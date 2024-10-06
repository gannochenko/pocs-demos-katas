package dto

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"api/internal/domain"
	"api/internal/util/db"
)

type Pet struct {
	gorm.Model
	ID         uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string           `db:"name"`
	Status     domain.PetStatus `db:"status"`
	PhotoUrls  []string         `gorm:"type:text[]"`
	CategoryID *uuid.UUID       `gorm:"type:uuid"`
	Tags       []Tag            `gorm:"many2many:pet_tags"`
	Category   *Category
}

func (p *Pet) ToDomain() (*domain.Pet, error) {
	result := &domain.Pet{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ListPetsFilter struct {
	ID     []string
	Status *domain.PetStatus
}

type ListPetParameters struct {
	Filter     *ListPetsFilter
	Pagination *db.Pagination
}

type ListTagsFilter struct {
	PetID *string
}

type ListTagsParameters struct {
	Filter     *ListTagsFilter
	Pagination *db.Pagination
}

type ListCategoriesFilter struct {
	PetID *string
}

type ListCategoriesParameters struct {
	Filter     *ListCategoriesFilter
	Pagination *db.Pagination
}
