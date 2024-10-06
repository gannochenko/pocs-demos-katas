package domain

type Category struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListCategoriesRequest struct {
	Pagination *PaginationRequest
}

type ListCategoriesResponse struct {
	Categories []*Category         `json:"categories"`
	Pagination *PaginationResponse `json:"pagination"`
}
