package interfaces

import (
	"context"

	"gorm.io/gorm"

	"api/internal/dto"
)

type PetRepository interface {
	ListPets(ctx context.Context, tx *gorm.DB, parameters *dto.ListPetParameters) (result []*dto.Pet, err error)
	CountPets(ctx context.Context, tx *gorm.DB, parameters *dto.ListPetParameters) (count int64, err error)
	UpdatePet(ctx context.Context, tx *gorm.DB, pet *dto.Pet) error
}

type TagRepository interface {
	ListTags(ctx context.Context, tx *gorm.DB, parameters *dto.ListTagsParameters) (result []*dto.Tag, err error)
	CountTags(ctx context.Context, tx *gorm.DB, parameters *dto.ListTagsParameters) (count int64, err error)
}

type CategoryRepository interface {
	ListCategories(ctx context.Context, tx *gorm.DB, parameters *dto.ListCategoriesParameters) (result []*dto.Category, err error)
	CountCategories(ctx context.Context, tx *gorm.DB, parameters *dto.ListCategoriesParameters) (count int64, err error)
}

type PetTagRepository interface {
}

type PetCategoryRepository interface {
}

type OrderRepository interface {
}

type CustomerRepository interface {
}

type AddressRepository interface {
}
