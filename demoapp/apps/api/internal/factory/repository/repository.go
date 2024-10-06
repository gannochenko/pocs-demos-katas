package repository

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/repository/category"
	"api/internal/repository/pet"
	"api/internal/repository/tag"
)

type Factory struct {
	session *gorm.DB

	petRepository         interfaces.PetRepository
	tagRepository         interfaces.TagRepository
	categoryRepository    interfaces.CategoryRepository
	petTagRepository      interfaces.PetTagRepository
	petCategoryRepository interfaces.PetCategoryRepository
	orderRepository       interfaces.OrderRepository
	customerRepository    interfaces.CustomerRepository
	addressRepository     interfaces.AddressRepository
}

func New(session *gorm.DB) *Factory {
	return &Factory{
		session: session,
	}
}

func (m *Factory) GetPetRepository() interfaces.PetRepository {
	if m.petRepository == nil {
		m.petRepository = pet.New(m.session)
	}

	return m.petRepository
}

func (m *Factory) GetTagRepository() interfaces.TagRepository {
	if m.tagRepository == nil {
		m.tagRepository = tag.New(m.session)
	}

	return m.tagRepository
}

func (m *Factory) GetCategoryRepository() interfaces.CategoryRepository {
	if m.categoryRepository == nil {
		m.categoryRepository = category.New(m.session)
	}

	return m.categoryRepository
}

func (m *Factory) GetPetTagRepository() interfaces.PetTagRepository {
	if m.petTagRepository == nil {
		m.petTagRepository = nil
	}

	return m.petTagRepository
}

func (m *Factory) GetPetCategoryRepository() interfaces.PetCategoryRepository {
	if m.petCategoryRepository == nil {
		m.petCategoryRepository = nil
	}

	return m.petCategoryRepository
}
