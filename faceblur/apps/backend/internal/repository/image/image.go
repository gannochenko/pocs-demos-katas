package image

import (
	"context"

	"gorm.io/gorm"

	"backend/internal/database"
)

type Repository struct {
	session *gorm.DB
}

func New(session *gorm.DB) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) Create(ctx context.Context, image *database.Image) error {
	res := r.session.WithContext(ctx).Create(image)
	return res.Error
}

func (r *Repository) List(ctx context.Context, parameters *database.ListParameters) ([]database.Image, error) {
	var result []database.Image
	res := r.session.WithContext(ctx).Where("age > ?", 30).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}

func (r *Repository) Count(ctx context.Context, parameters *database.CountParameters) (int64, error) {
	var result int64
	err := r.session.WithContext(ctx).Model(&database.Image{}).Where("age > ?", 30).Count(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
