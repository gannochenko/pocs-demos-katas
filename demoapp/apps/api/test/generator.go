package test

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"

	"api/internal/domain"
	"api/internal/dto"
)

func NewGenerator() *Generator {
	return &Generator{}
}

type Generator struct {
}

func (g *Generator) CreateUUID() uuid.UUID {
	return uuid.New()
}

func (g *Generator) CreatePet() *dto.Pet {
	return &dto.Pet{
		ID:         g.CreateUUID(),
		Name:       gofakeit.Name(),
		Status:     domain.PetStatusAvailable,
		CategoryID: nil,
	}
}

func (g *Generator) CreateCategory() *dto.Category {
	return &dto.Category{
		ID:   g.CreateUUID(),
		Name: gofakeit.BeerName(),
	}
}

func (g *Generator) CreateTag() *dto.Tag {
	return &dto.Tag{
		ID:   g.CreateUUID(),
		Name: gofakeit.Dog(),
	}
}

func (g *Generator) CreatePetTag() *dto.PetTag {
	return &dto.PetTag{
		PetID: g.CreateUUID(),
		TagID: g.CreateUUID(),
	}
}

func (g *Generator) CreateOrder() *dto.Order {
	return &dto.Order{
		ID:       g.CreateUUID(),
		PetID:    g.CreateUUID(),
		Quantity: g.CreatePositiveInt32(),
		ShipDate: g.CreateUTCTime(),
		Status:   domain.OrderStatusPlaced,
		Complete: true,
	}
}

func (g *Generator) CreateCustomer() *dto.Customer {
	return &dto.Customer{
		ID:       g.CreateUUID(),
		Username: gofakeit.Email(),
	}
}

func (g *Generator) CreateAddress() *dto.Address {
	return &dto.Address{
		ID:     g.CreateUUID(),
		Street: gofakeit.Street(),
		City:   gofakeit.City(),
		State:  gofakeit.State(),
		Zip:    gofakeit.Zip(),
	}
}

func (g *Generator) CreatePositiveInt32() int32 {
	return int32(gofakeit.Uint16() + 1)
}

func (g *Generator) CreateUTCTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}
