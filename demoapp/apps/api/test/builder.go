package test

import (
	"gorm.io/gorm"

	"api/internal/dto"
)

type Builder struct {
	session *gorm.DB

	pets          []*dto.Pet
	petCategories []*dto.PetCategory
	petTags       []*dto.PetTag
	customers     []*dto.Customer
	addresses     []*dto.Address
	orders        []*dto.Order
}

func NewBuilder(session *gorm.DB) *Builder {
	writer := &Builder{
		session: session,
	}
	writer.resetArrays()

	return writer
}

func (b *Builder) AddPet(pet *dto.Pet) {
	b.pets = append(b.pets, pet)
}

func (b *Builder) AddPets(pets ...*dto.Pet) *Builder {
	for _, pet := range pets {
		b.AddPet(pet)
	}

	return b
}

func (b *Builder) AddPetTag(petTag *dto.PetTag) {
	b.petTags = append(b.petTags, petTag)
}

func (b *Builder) AddPetTags(petTags ...*dto.PetTag) *Builder {
	for _, petTag := range petTags {
		b.AddPetTag(petTag)
	}

	return b
}

func (b *Builder) AddPetCategory(petCategory *dto.PetCategory) {
	b.petCategories = append(b.petCategories, petCategory)
}

func (b *Builder) AddPetCategories(petCategories ...*dto.PetCategory) *Builder {
	for _, petCategory := range petCategories {
		b.AddPetCategory(petCategory)
	}

	return b
}

func (b *Builder) AddCustomer(customer *dto.Customer) {
	b.customers = append(b.customers, customer)
}

func (b *Builder) AddCustomers(customers ...*dto.Customer) *Builder {
	for _, customer := range customers {
		b.AddCustomer(customer)
	}

	return b
}

func (b *Builder) AddOrder(order *dto.Order) {
	b.orders = append(b.orders, order)
}

func (b *Builder) AddOrders(orders ...*dto.Order) *Builder {
	for _, order := range orders {
		b.AddOrder(order)
	}

	return b
}

func (b *Builder) AddAddress(address *dto.Address) {
	b.addresses = append(b.addresses, address)
}

func (b *Builder) AddAddresses(addresses ...*dto.Address) *Builder {
	for _, address := range addresses {
		b.AddAddress(address)
	}

	return b
}

func (b *Builder) Submit() error {
	for _, pet := range b.pets {
		res := b.session.Create(pet)
		if res.Error != nil {
			return res.Error
		}
	}

	for _, petTag := range b.petTags {
		res := b.session.Create(petTag)
		if res.Error != nil {
			return res.Error
		}
	}

	for _, petCategory := range b.petCategories {
		res := b.session.Create(petCategory)
		if res.Error != nil {
			return res.Error
		}
	}

	for _, customer := range b.customers {
		res := b.session.Create(customer)
		if res.Error != nil {
			return res.Error
		}
	}

	for _, order := range b.orders {
		res := b.session.Create(order)
		if res.Error != nil {
			return res.Error
		}
	}

	for _, address := range b.addresses {
		res := b.session.Create(address)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (b *Builder) Reset() *Builder {
	b.resetArrays()

	return b
}

func (b *Builder) SelectPets(filter string) (result []*dto.Pet, err error) {
	queryResult := b.session.Table("pets").Where(filter).Find(&result)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	return result, nil
}

func (b *Builder) resetArrays() {
	b.pets = make([]*dto.Pet, 0)
	b.petCategories = make([]*dto.PetCategory, 0)
	b.petTags = make([]*dto.PetTag, 0)
	b.customers = make([]*dto.Customer, 0)
	b.orders = make([]*dto.Order, 0)
	b.addresses = make([]*dto.Address, 0)
}
