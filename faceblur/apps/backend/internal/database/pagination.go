package database

import "backend/internal/domain"

type Pagination struct {
	PageNumber int32
	PageSize   int32
}

func NewPaginationFromDomainRequest(pagination *domain.PaginationRequest) *Pagination {
	return &Pagination{
		PageSize:   pagination.PageSize,
		PageNumber: pagination.PageNumber,
	}
}
