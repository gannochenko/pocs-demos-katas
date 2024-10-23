package factory

import (
	"gorm.io/gorm"

	"api/interfaces"
	"api/internal/factory/repository"
	"api/internal/factory/service"
)

func MakeServiceFactory(session *gorm.DB, configService interfaces.ConfigService) *service.Factory {
	return service.New(session, repository.New(session), configService)
}
