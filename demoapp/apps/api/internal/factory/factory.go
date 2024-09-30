package factory

import (
	"gorm.io/gorm"

	"api/internal/factory/repository"
	"api/internal/factory/service"
)

func MakeServiceFactory(session *gorm.DB) *service.Factory {
	return service.New(session, repository.New(session))
}
