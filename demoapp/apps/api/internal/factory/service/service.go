package service

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/factory/repository"
	"api/internal/service"
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
		m.petService = service.NewPetService()
	}

	return m.petService
}

func (m *Factory) GetStoreService() interfaces.StoreService {
	if m.storeService == nil {
		m.storeService = service.NewStoreService()
	}

	return m.storeService
}
