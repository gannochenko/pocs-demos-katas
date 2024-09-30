package pet

import (
	"context"

	"api/internal/domain"
	"api/pkg/syserr"
)

// Service is a service that implements the logic for the PetAPIServicer
// This service should implement the business logic for every endpoint for the PetAPI API.
// Include any external packages or services that will be required by this service.
type Service struct {
}

// NewPetService creates a default api service
func NewPetService() *Service {
	return &Service{}
}

// AddPet - Add a new pet to the store
func (s *Service) AddPet(ctx context.Context, pet *domain.Pet) (*domain.AddPetResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// UpdatePet - Update an existing pet
func (s *Service) UpdatePet(ctx context.Context, pet *domain.Pet) (*domain.UpdatePetResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// DeletePet - Deletes a pet
func (s *Service) DeletePet(ctx context.Context, petId int64, apiKey string) (*domain.DeletePetResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// ListPets - Finds Pets by filter
func (s *Service) ListPets(ctx context.Context, status string) (*domain.ListPetsResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// GetPetByID - Find pet by ID
func (s *Service) GetPetByID(ctx context.Context, petId int64) (*domain.GetPetByIdResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}
