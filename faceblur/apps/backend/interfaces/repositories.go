package interfaces

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/database"
)

type UserRepository interface {
	List(ctx context.Context, tx *gorm.DB, parameters database.UserListParameters) ([]database.User, error)
}

type ImageRepository interface {
	Create(ctx context.Context, tx *gorm.DB, image *database.Image) error
	Update(ctx context.Context, tx *gorm.DB, image *database.ImageUpdate) error
	List(ctx context.Context, tx *gorm.DB, parameters database.ImageListParameters) ([]database.Image, error)
	GetByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*database.Image, error)
	Count(ctx context.Context, tx *gorm.DB, parameters database.ImageCountParameters) (int32, error)
}

type ImageProcessingQueueRepository interface {
	Create(ctx context.Context, tx *gorm.DB, element *database.ImageProcessingQueue) error
	Update(ctx context.Context, tx *gorm.DB, element *database.ImageProcessingQueueUpdate) error
	List(ctx context.Context, tx *gorm.DB, parameters database.ImageProcessingQueueListParameters) ([]database.ImageProcessingQueue, error)
}
