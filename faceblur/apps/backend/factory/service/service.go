package service

import (
	"io"

	"gorm.io/gorm"

	"backend/factory/repository"
	"backend/interfaces"
	"backend/internal/service/auth"
	"backend/internal/service/config"
	"backend/internal/service/eventBus"
	"backend/internal/service/facedetection"
	"backend/internal/service/image"
	"backend/internal/service/imageProcessor"
	"backend/internal/service/logger"
	"backend/internal/service/monitoring"
	"backend/internal/service/storage"
	"backend/internal/service/user"
	"backend/internal/util/db"
)

type Factory struct {
	session           *gorm.DB
	outputWriter      io.Writer
	repositoryFactory *repository.Factory

	sessionManager interfaces.SessionManager

	configService         interfaces.ConfigService
	loggerService         interfaces.LoggerService
	imageService          interfaces.ImageService
	storageService        interfaces.StorageService
	authService           interfaces.AuthService
	userService           interfaces.UserService
	eventBusService       interfaces.EventBusService
	imageProcessorService interfaces.ImageProcessorService
	faceDetectionService interfaces.FaceDetectionService
	monitoringService interfaces.MonitoringService
}

func NewServiceFactory(session *gorm.DB, outputWriter io.Writer, repositoryFactory *repository.Factory) *Factory {
	return &Factory{
		session:           session,
		repositoryFactory: repositoryFactory,
		outputWriter:      outputWriter,
	}
}

func (f *Factory) GetRepositoryFactory() *repository.Factory {
	return f.repositoryFactory
}

func (f *Factory) GetSessionManager() interfaces.SessionManager {
	if f.sessionManager == nil {
		f.sessionManager = db.NewSessionManager(f.session)
	}

	return f.sessionManager
}

func (f *Factory) GetConfigService() interfaces.ConfigService {
	if f.configService == nil {
		f.configService = config.NewConfigService()
	}

	return f.configService
}

func (f *Factory) GetLoggerService() interfaces.LoggerService {
	if f.loggerService == nil {
		f.loggerService = logger.NewLoggerService(f.outputWriter)
	}

	return f.loggerService
}

func (f *Factory) GetImageService() interfaces.ImageService {
	if f.imageService == nil {
		f.imageService = image.NewImageService(
			f.GetSessionManager(),
			f.repositoryFactory.GetImageRepository(),
			f.repositoryFactory.GetImageProcessingQueueRepository(),
			f.GetStorageService(),
			f.GetConfigService(),
			f.GetEventBusService(),
			f.GetLoggerService(),
		)
	}

	return f.imageService
}

func (f *Factory) GetStorageService() interfaces.StorageService {
	if f.storageService == nil {
		f.storageService = storage.NewStorageService(
			f.GetConfigService(),
			f.GetLoggerService(),
		)
	}

	return f.storageService
}

func (f *Factory) GetAuthService() interfaces.AuthService {
	if f.authService == nil {
		f.authService = auth.NewAuthService(
			f.GetConfigService(),
			f.GetLoggerService(),
		)
	}

	return f.authService
}

func (f *Factory) GetUserService() interfaces.UserService {
	if f.userService == nil {
		f.userService = user.NewUserService(
			f.GetRepositoryFactory().GetUserRepository(),
		)
	}

	return f.userService
}

func (f *Factory) GetEventBusService() interfaces.EventBusService {
	if f.eventBusService == nil {
		f.eventBusService = eventBus.NewEventBusService(f.GetConfigService(), f.GetLoggerService())
	}

	return f.eventBusService
}

func (f *Factory) GetImageProcessorService() interfaces.ImageProcessorService {
	if f.imageProcessorService == nil {
		f.imageProcessorService = imageProcessor.NewImageProcessor(
			f.GetConfigService(),
			f.GetEventBusService(),
			f.GetLoggerService(),
			f.GetRepositoryFactory().GetImageProcessingQueueRepository(),
			f.GetRepositoryFactory().GetImageRepository(),
			f.GetFaceDetectionService(),
			f.GetStorageService(),
		)
	}

	return f.imageProcessorService
}

func (f *Factory) GetFaceDetectionService() interfaces.FaceDetectionService {
	if f.faceDetectionService == nil {
		f.faceDetectionService = facedetection.NewService(
			f.GetConfigService(),
			f.GetLoggerService(),
		)
	}

	return f.faceDetectionService
}

func (f *Factory) GetMonitoringService() interfaces.MonitoringService {
	if f.monitoringService == nil {
		f.monitoringService = monitoring.NewService(f.GetConfigService())
	}

	return f.monitoringService
}
