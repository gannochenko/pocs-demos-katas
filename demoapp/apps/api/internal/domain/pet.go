package domain

type PetCategory struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PetTag struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

type Pet struct {
	Id        int64       `json:"id,omitempty"`
	Name      string      `json:"name"`
	Category  PetCategory `json:"category,omitempty"`
	PhotoUrls []string    `json:"photoUrls"`
	Tags      []PetTag    `json:"tags,omitempty"`
	Status    PetStatus   `json:"status,omitempty"`
}

type AddPetResponse struct{}

type DeletePetResponse struct{}

type ListPetsResponse struct{}

type FindPetsByTagsResponse struct{}

type GetPetByIdResponse struct{}

type UpdatePetResponse struct{}

type UpdatePetWithFormResponse struct{}

type UploadFileResponse struct{}
