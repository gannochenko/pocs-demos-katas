package test

import (
	"gorm.io/gorm"

	"api/internal/dto"
)

type Writer struct {
	session *gorm.DB

	pets          []*dto.Pet
	petCategories []*dto.PetCategory
	petTags       []*dto.PetTag
	customers     []*dto.Customer
	addresses     []*dto.Address
	orders        []*dto.Order
}

func NewBuilder(session *gorm.DB) *Writer {
	writer := &Writer{
		session: session,
	}
	writer.resetArrays()

	return writer
}

func (w *Writer) AddPet(pet *dto.Pet) {
	w.pets = append(w.pets, pet)
}

func (w *Writer) AddPets(pets ...*dto.Pet) *Writer {
	for _, pet := range pets {
		w.AddPet(pet)
	}

	return w
}

func (w *Writer) AddPetTag(petTag *dto.PetTag) {
	w.petTags = append(w.petTags, petTag)
}

func (w *Writer) AddPetTags(petTags ...*dto.PetTag) *Writer {
	for _, petTag := range petTags {
		w.AddPetTag(petTag)
	}

	return w
}

func (w *Writer) AddPetCategory(petCategory *dto.PetCategory) {
	w.petCategories = append(w.petCategories, petCategory)
}

func (w *Writer) AddPetCategories(petCategories ...*dto.PetCategory) *Writer {
	for _, petCategory := range petCategories {
		w.AddPetCategory(petCategory)
	}

	return w
}

func (w *Writer) AddCustomer(customer *dto.Customer) {
	w.customers = append(w.customers, customer)
}

func (w *Writer) AddCustomers(customers ...*dto.Customer) *Writer {
	for _, customer := range customers {
		w.AddCustomer(customer)
	}

	return w
}

func (w *Writer) AddOrder(order *dto.Order) {
	w.orders = append(w.orders, order)
}

func (w *Writer) AddOrders(orders ...*dto.Order) *Writer {
	for _, order := range orders {
		w.AddOrder(order)
	}

	return w
}

func (w *Writer) AddAddress(address *dto.Address) {
	w.addresses = append(w.addresses, address)
}

func (w *Writer) AddAddresses(addresses ...*dto.Address) *Writer {
	for _, address := range addresses {
		w.AddAddress(address)
	}

	return w
}

func (w *Writer) Submit() error {
	for _, pet := range w.pets {
		w.session.Create(pet)
	}

	for _, petTag := range w.petTags {
		w.session.Create(petTag)
	}

	for _, petCategory := range w.petCategories {
		w.session.Create(petCategory)
	}

	for _, customer := range w.customers {
		w.session.Create(customer)
	}

	for _, order := range w.orders {
		w.session.Create(order)
	}

	for _, address := range w.addresses {
		w.session.Create(address)
	}

	return nil
}

func (w *Writer) Reset() *Writer {
	w.resetArrays()

	return w
}

func (w *Writer) resetArrays() {
	w.pets = make([]*dto.Pet, 0)
	w.petCategories = make([]*dto.PetCategory, 0)
	w.petTags = make([]*dto.PetTag, 0)
	w.customers = make([]*dto.Customer, 0)
	w.orders = make([]*dto.Order, 0)
	w.addresses = make([]*dto.Address, 0)
}
