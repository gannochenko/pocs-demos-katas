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

func (f *Factory) GetRepositoryFactory() *repository.Factory {
	return f.repositoryFactory
}

func (f *Factory) GetPetService() interfaces.PetService {
	if f.petService == nil {
		f.petService = pet.NewPetService()
	}

	return f.petService
}

func (f *Factory) GetStoreService() interfaces.StoreService {
	if f.storeService == nil {
		f.storeService = store.NewStoreService()
	}

	return f.storeService
}
