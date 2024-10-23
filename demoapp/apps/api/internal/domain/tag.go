package domain

type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListTagsRequest struct {
	Pagination *PaginationRequest
}

type ListTagsResponse struct {
	Pets       []*Tag              `json:"tags"`
	Pagination *PaginationResponse `json:"pagination"`
}
