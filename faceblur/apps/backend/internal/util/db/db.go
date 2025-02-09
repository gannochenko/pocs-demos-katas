package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/database"
)

func Connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func GetRunner(session *gorm.DB, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}

	return session
}

func ApplyPagination(runner *gorm.DB, pagination *database.Pagination) *gorm.DB {
	if pagination != nil {
		runner = runner.Offset(int((pagination.PageNumber - 1) * pagination.PageSize)).Limit(int(pagination.PageSize))
	}

	return runner
}
