package domain

import (
	"fmt"

	"api/pkg/syserr"
)

type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

type Pet struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  *Category `json:"category,omitempty"`
	PhotoUrls []string  `json:"photoUrls"`
	Tags      []Tag     `json:"tags"`
	Status    PetStatus `json:"status"`
}

func (p *Pet) IsValid() error {
	var errors []string

	// todo: support i18n here
	if p.ID == "" {
		errors = append(errors, "id not set")
	}
	if p.Name == "" {
		errors = append(errors, "name not set")
	}
	if p.Category.ID == "" {
		errors = append(errors, "category not set")
	}
	if p.Status == "" {
		errors = append(errors, "status not set")
	} else {
		if p.Status != PetStatusSold && p.Status != PetStatusAvailable && p.Status != PetStatusPending {
			errors = append(errors, fmt.Sprintf("unknown status: %s", p.Status))
		}
	}

	if len(errors) > 0 {
		return syserr.NewBadInput("validation failed", syserr.F("reasons", errors))
	}

	return nil
}

type ListPetsRequest struct {
	Status     PetStatus
	IDs        []string
	Pagination *PaginationRequest
}

type ListPetsResponse struct {
	Pets       []*Pet              `json:"pets"`
	Pagination *PaginationResponse `json:"pagination"`
}

type GetPetRequest struct {
	ID string `json:"id"`
}

type GetPetResponse struct {
	Pet *Pet `json:"pet"`
}

type UpdatePetRequest struct {
	Pet *Pet `json:"pet"`
}

type UpdatePetResponse struct {
}

type AddPetRequest struct {
	Pet *Pet `json:"pet"`
}

type AddPetResponse struct {
}

type DeletePetRequest struct {
	ID string `json:"id"`
}

type DeletePetResponse struct{}
