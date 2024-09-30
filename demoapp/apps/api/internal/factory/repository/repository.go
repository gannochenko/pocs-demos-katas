package repository

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/repository/pet"
)

type Factory struct {
	session *gorm.DB

	petRepository         interfaces.PetRepository
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

func (m *Factory) GetPetTagRepository() interfaces.PetTagRepository {
	if m.petTagRepository == nil {
		m.petTagRepository = nil
	}

	return m.petTagRepository
}
