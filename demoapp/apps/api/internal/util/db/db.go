package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	PageNumber int
	PageSize   int
}

func ApplyPagination(runner *gorm.DB, pagination *Pagination) *gorm.DB {
	if pagination != nil {
		runner = runner.Offset(int(pagination.PageNumber * pagination.PageSize)).Limit(pagination.PageSize)
	}

	return runner
}
