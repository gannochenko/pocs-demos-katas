package service

import (
	"context"
	"os"

	"api/internal/domain"
	"api/pkg/syserr"
)

// PetAPIService is a service that implements the logic for the PetAPIServicer
// This service should implement the business logic for every endpoint for the PetAPI API.
// Include any external packages or services that will be required by this service.
type PetAPIService struct {
}

// NewPetAPIService creates a default api service
func NewPetAPIService() *PetAPIService {
	return &PetAPIService{}
}

// AddPet - Add a new pet to the store
func (s *PetAPIService) AddPet(ctx context.Context, pet *domain.Pet) (*domain.AddPetResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// DeletePet - Deletes a pet
func (s *PetAPIService) DeletePet(ctx context.Context, petId int64, apiKey string) (*domain.DeletePetResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// FindPetsByStatus - Finds Pets by status
func (s *PetAPIService) FindPetsByStatus(ctx context.Context, status string) (*domain.FindPetsByStatusResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// FindPetsByTags - Finds Pets by tags
func (s *PetAPIService) FindPetsByTags(ctx context.Context, tags []string) (*domain.FindPetsByTagsResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// GetPetById - Find pet by ID
func (s *PetAPIService) GetPetById(ctx context.Context, petId int64) (*domain.GetPetByIdResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// UpdatePet - Update an existing pet
func (s *PetAPIService) UpdatePet(ctx context.Context, pet *domain.Pet) (*domain.UpdatePetResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (s *PetAPIService) UpdatePetWithForm(ctx context.Context, petId int64, name string, status string) (*domain.UpdatePetWithFormResponse, error) {
	return nil, syserr.NewNotImplemented("")
}

// UploadFile - uploads an image
func (s *PetAPIService) UploadFile(ctx context.Context, petId int64, additionalMetadata string, body *os.File) (*domain.UploadFileResponse, error) {
	return nil, syserr.NewNotImplemented("")
}
