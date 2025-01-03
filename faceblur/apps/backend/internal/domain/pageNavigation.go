package domain

import "math"

type PageNavigationRequest struct {
	PageNumber int32
	PageSize   int32
}

type PageNavigationResponse struct {
	PageNumber int32 `json:"page_number"`
	PageSize   int32 `json:"page_size"`
	PageCount  int32 `json:"page_count"`
	Total      int32 `json:"total"`
}

func NewPageNavigationResponseFromRequest(request *PageNavigationRequest, total int32) *PageNavigationResponse {
	if request == nil {
		return &PageNavigationResponse{
			PageNumber: 1,
			PageSize:   math.MaxInt32,
			Total:      total,
			PageCount:  1,
		}
	}

	return &PageNavigationResponse{
		PageNumber: request.PageNumber,
		PageSize:   request.PageSize,
		Total:      total,
		PageCount:  int32(math.Ceil(float64(total) / float64(request.PageSize))),
	}
}
