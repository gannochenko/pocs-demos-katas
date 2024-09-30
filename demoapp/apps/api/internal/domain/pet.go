package domain

type Category struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Tag struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Pet struct {
	Id        int64    `json:"id,omitempty"`
	Name      string   `json:"name"`
	Category  Category `json:"category,omitempty"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags,omitempty"`
	Status    string   `json:"status,omitempty"`
}

type AddPetResponse struct{}

type DeletePetResponse struct{}

type FindPetsByStatusResponse struct{}

type FindPetsByTagsResponse struct{}

type GetPetByIdResponse struct{}

type UpdatePetResponse struct{}

type UpdatePetWithFormResponse struct{}

type UploadFileResponse struct{}
