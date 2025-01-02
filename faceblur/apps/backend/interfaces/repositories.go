package interfaces

import (
	"context"

	"gorm.io/gorm"

	"backend/internal/database"
)

type UserRepository interface {
}

type ImageRepository interface {
	Create(ctx context.Context, tx *gorm.DB, image *database.Image) error
	Update(ctx context.Context, tx *gorm.DB, image *database.Image) error
	List(ctx context.Context, tx *gorm.DB, parameters database.ImageListParameters) ([]database.Image, error)
	Count(ctx context.Context, tx *gorm.DB, parameters database.ImageCountParameters) (int64, error)
}

type ImageProcessingQueueRepository interface {
	Create(ctx context.Context, tx *gorm.DB, element *database.ImageProcessingQueue) error
	Update(ctx context.Context, tx *gorm.DB, element *database.ImageProcessingQueue) error
	List(ctx context.Context, tx *gorm.DB, parameters database.ImageProcessingQueueListParameters) ([]database.ImageProcessingQueue, error)
}
