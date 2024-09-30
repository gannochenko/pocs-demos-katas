package service

import (
	"context"
	"os"

	"api/internal/domain"
	"api/pkg/syserr"
)

// PetService is a service that implements the logic for the PetAPIServicer
// This service should implement the business logic for every endpoint for the PetAPI API.
// Include any external packages or services that will be required by this service.
type PetService struct {
}

// NewPetService creates a default api service
func NewPetService() *PetService {
	return &PetService{}
}

// AddPet - Add a new pet to the store
func (s *PetService) AddPet(ctx context.Context, pet *domain.Pet) (*domain.AddPetResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// DeletePet - Deletes a pet
func (s *PetService) DeletePet(ctx context.Context, petId int64, apiKey string) (*domain.DeletePetResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// FindPetsByStatus - Finds Pets by status
func (s *PetService) FindPetsByStatus(ctx context.Context, status string) (*domain.FindPetsByStatusResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// FindPetsByTags - Finds Pets by tags
func (s *PetService) FindPetsByTags(ctx context.Context, tags []string) (*domain.FindPetsByTagsResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// GetPetById - Find pet by ID
func (s *PetService) GetPetById(ctx context.Context, petId int64) (*domain.GetPetByIdResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// UpdatePet - Update an existing pet
func (s *PetService) UpdatePet(ctx context.Context, pet *domain.Pet) (*domain.UpdatePetResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (s *PetService) UpdatePetWithForm(ctx context.Context, petId int64, name string, status string) (*domain.UpdatePetWithFormResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}

// UploadFile - uploads an image
func (s *PetService) UploadFile(ctx context.Context, petId int64, additionalMetadata string, body *os.File) (*domain.UploadFileResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}
