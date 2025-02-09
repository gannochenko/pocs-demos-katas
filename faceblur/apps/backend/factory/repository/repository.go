package repository

import (
	"gorm.io/gorm"

	"backend/interfaces"
	"backend/internal/repository/image"
	"backend/internal/repository/imageProcessingQueue"
	"backend/internal/repository/user"
)

type Factory struct {
	session *gorm.DB

	userRepository                 interfaces.UserRepository
	imageRepository                interfaces.ImageRepository
	imageProcessingQueueRepository interfaces.ImageProcessingQueueRepository
}

func NewRepositoryFactory(session *gorm.DB) *Factory {
	return &Factory{
		session: session,
	}
}

func (f *Factory) GetUserRepository() interfaces.UserRepository {
	if f.userRepository == nil {
		f.userRepository = user.NewUserRepository(f.session)
	}

	return f.userRepository
}

func (f *Factory) GetImageRepository() interfaces.ImageRepository {
	if f.imageRepository == nil {
		f.imageRepository = image.NewImageRepository(f.session)
	}

	return f.imageRepository
}

func (f *Factory) GetImageProcessingQueueRepository() interfaces.ImageProcessingQueueRepository {
	if f.imageProcessingQueueRepository == nil {
		f.imageProcessingQueueRepository = imageProcessingQueue.NewImageProcessingQueueRepository(f.session)
	}

	return f.imageProcessingQueueRepository
}
