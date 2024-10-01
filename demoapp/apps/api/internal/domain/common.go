package domain

import "math"

const (
	MaxPageSize     = 200
	DefaultPageSize = 50
)

type PaginationRequest struct {
	PageNumber int32
	PageSize   int32
}

func SanitizePaginationRequest(request *PaginationRequest) *PaginationRequest {
	if request == nil {
		return &PaginationRequest{
			PageNumber: 1,
			PageSize:   DefaultPageSize,
		}
	}

	if request.PageSize <= 0 {
		request.PageSize = DefaultPageSize
	}
	if request.PageNumber <= 0 {
		request.PageNumber = 1
	}

	return request
}

type PaginationResponse struct {
	PageNumber int32 `json:"page_number"`
	PageSize   int32 `json:"page_size"`
	PageCount  int64 `json:"page_count"`
	Total      int64 `json:"total"`
}

func NewPaginationResponseFromRequest(request *PaginationRequest, total int64) *PaginationResponse {
	return &PaginationResponse{
		PageNumber: request.PageNumber,
		PageSize:   request.PageSize,
		Total:      total,
		PageCount:  int64(math.Ceil(float64(total) / float64(request.PageSize))),
	}
}
