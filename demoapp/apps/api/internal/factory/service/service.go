package service

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/factory/repository"
	"api/internal/service/auth"
	"api/internal/service/category"
	"api/internal/service/pet"
	"api/internal/service/store"
	"api/internal/service/tag"
)

type Factory struct {
	session           *gorm.DB
	repositoryFactory *repository.Factory

	configService   interfaces.ConfigService
	petService      interfaces.PetService
	tagService      interfaces.TagService
	categoryService interfaces.CategoryService
	storeService    interfaces.StoreService
	authService     interfaces.AuthService
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

func (f *Factory) SetPetService(foo interfaces.PetService) {
	f.petService = foo
}

func (f *Factory) GetPetService() interfaces.PetService {
	if f.petService == nil {
		f.petService = pet.New(
			f.repositoryFactory.GetPetRepository(),
		)
	}

	return f.petService
}

func (f *Factory) GetTagService() interfaces.TagService {
	if f.tagService == nil {
		f.tagService = tag.New(
			f.repositoryFactory.GetTagRepository(),
		)
	}

	return f.tagService
}

func (f *Factory) GetCategoryService() interfaces.CategoryService {
	if f.categoryService == nil {
		f.categoryService = category.New(
			f.repositoryFactory.GetCategoryRepository(),
		)
	}

	return f.categoryService
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
