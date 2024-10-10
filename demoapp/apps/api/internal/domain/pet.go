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
