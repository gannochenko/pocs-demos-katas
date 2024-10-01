package pet

import (
	"context"

	"api/interfaces"
	"api/internal/domain"
	"api/internal/dto"
	"api/internal/util/db"
	"api/pkg/syserr"
)

// Service is a service that implements the logic for the PetAPIServicer
// This service should implement the business logic for every endpoint for the PetAPI API.
// Include any external packages or services that will be required by this service.
type Service struct {
	petRepository         interfaces.PetRepository
	petTagRepository      interfaces.PetTagRepository
	petCategoryRepository interfaces.PetCategoryRepository
}

// NewPetService creates a default api service
func NewPetService(
	petRepository interfaces.PetRepository,
	petTagRepository interfaces.PetTagRepository,
	petCategoryRepository interfaces.PetCategoryRepository,
) *Service {
	return &Service{
		petRepository:         petRepository,
		petTagRepository:      petTagRepository,
		petCategoryRepository: petCategoryRepository,
	}
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
func (s *Service) ListPets(ctx context.Context, request *domain.ListPetsRequest) (*domain.ListPetsResponse, error) {
	request.Pagination = domain.SanitizePaginationRequest(request.Pagination)

	filter := &dto.ListPetFilter{}
	if request.Status != "" {
		filter.Status = &request.Status
	}
	if len(request.IDs) > 0 {
		filter.ID = request.IDs
	}

	count, err := s.petRepository.CountPets(ctx, nil, &dto.ListPetParameters{
		Filter: filter,
	})
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not count pets")
	}

	result := &domain.ListPetsResponse{
		Pagination: domain.NewPaginationResponseFromRequest(request.Pagination, count),
	}

	if count > 0 {
		res, err := s.petRepository.ListPets(ctx, nil, &dto.ListPetParameters{
			Filter:     filter,
			Pagination: db.NewPaginationFromDomainRequest(request.Pagination),
		})
		if err != nil {
			return nil, syserr.Wrap(err, syserr.InternalCode, "could not list pets")
		}

		for _, pet := range res {
			domainPet, err := pet.ToDomain()
			if err != nil {
				return nil, syserr.Wrap(err, syserr.InternalCode, "could not convert pet to domain")
			}
			result.Pets = append(result.Pets, domainPet)
		}
	}

	return result, nil
}

// GetPetByID - Find pet by ID
func (s *Service) GetPetByID(ctx context.Context, petId int64) (*domain.GetPetByIdResponse, error) {
	return nil, syserr.NewNotImplemented("method not implemented")
}
