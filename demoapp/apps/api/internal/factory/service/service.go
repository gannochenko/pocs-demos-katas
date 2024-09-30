package service

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/factory/repository"
	"api/internal/service/pet"
	"api/internal/service/store"
)

type Factory struct {
	session           *gorm.DB
	repositoryFactory *repository.Factory

	petService   interfaces.PetService
	storeService interfaces.StoreService
}

func New(session *gorm.DB, repositoryFactory *repository.Factory) *Factory {
	return &Factory{
		session:           session,
		repositoryFactory: repositoryFactory,
	}
}

func (m *Factory) GetPetService() interfaces.PetService {
	if m.petService == nil {
		m.petService = pet.NewPetService()
	}

	return m.petService
}

func (m *Factory) GetStoreService() interfaces.StoreService {
	if m.storeService == nil {
		m.storeService = store.NewStoreService()
	}

	return m.storeService
}
