package service

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/factory/repository"
	"api/internal/service/auth"
	"api/internal/service/pet"
	"api/internal/service/store"
)

type Factory struct {
	session           *gorm.DB
	repositoryFactory *repository.Factory

	configService interfaces.ConfigService
	petService    interfaces.PetService
	storeService  interfaces.StoreService
	authService   interfaces.AuthService
}

func New(session *gorm.DB, repositoryFactory *repository.Factory, configService interfaces.ConfigService) *Factory {
	return &Factory{
		session:           session,
		repositoryFactory: repositoryFactory,
		configService:     configService,
	}
}

func (f *Factory) GetRepositoryFactory() *repository.Factory {
	return f.repositoryFactory
}

func (f *Factory) GetPetService() interfaces.PetService {
	if f.petService == nil {
		f.petService = pet.NewPetService(
			f.repositoryFactory.GetPetRepository(),
			f.repositoryFactory.GetPetTagRepository(),
			f.repositoryFactory.GetPetCategoryRepository(),
		)
	}

	return f.petService
}

func (f *Factory) GetStoreService() interfaces.StoreService {
	if f.storeService == nil {
		f.storeService = store.NewStoreService()
	}

	return f.storeService
}

func (f *Factory) GetAuthService() interfaces.AuthService {
	if f.authService == nil {
		f.authService = auth.New(f.configService)
	}

	return f.authService
}
