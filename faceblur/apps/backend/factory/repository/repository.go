package repository

import (
	"gorm.io/gorm"

	"backend/interfaces"
	"backend/internal/repository/image"
)

type Factory struct {
	session *gorm.DB

	imageRepository interfaces.ImageRepository
}

func NewRepositoryFactory(session *gorm.DB) *Factory {
	return &Factory{
		session: session,
	}
}

func (f *Factory) GetImageRepository() interfaces.ImageRepository {
	if f.imageRepository == nil {
		f.imageRepository = image.NewImageRepository(f.session)
	}

	return f.imageRepository
}
