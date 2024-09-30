package interfaces

import (
	"context"

	"gorm.io/gorm"

	"api/internal/dto"
)

type PetRepository interface {
	ListPets(ctx context.Context, tx *gorm.DB, parameters *dto.ListPetParameters) (result []*dto.Pet, err error)
	CountPets(ctx context.Context, tx *gorm.DB, parameters *dto.ListPetParameters) (count int64, err error)
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
