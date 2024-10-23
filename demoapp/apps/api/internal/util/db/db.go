package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/internal/domain"
)

func Connect(dsn string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func GetRunner(session *gorm.DB, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return session
}

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

func ApplyPagination(runner *gorm.DB, pagination *Pagination) *gorm.DB {
	if pagination != nil {
		runner = runner.Offset(int((pagination.PageNumber - 1) * pagination.PageSize)).Limit(int(pagination.PageSize))
	}

	return runner
}
