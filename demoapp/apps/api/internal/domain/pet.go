package domain

type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

type Pet struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Category  Category  `json:"category,omitempty"`
	PhotoUrls []string  `json:"photoUrls"`
	Tags      []Tag     `json:"tags,omitempty"`
	Status    PetStatus `json:"status,omitempty"`
}

type AddPetResponse struct{}

type DeletePetResponse struct{}

type ListPetsRequest struct {
	Status     PetStatus
	IDs        []string
	Pagination *PaginationRequest
}

type ListPetsResponse struct {
	Pets       []*Pet              `json:"pets"`
	Pagination *PaginationResponse `json:"pagination"`
}

type FindPetsByTagsResponse struct{}

type GetPetByIdResponse struct{}

type UpdatePetResponse struct{}

type UpdatePetWithFormResponse struct{}

type UploadFileResponse struct{}
