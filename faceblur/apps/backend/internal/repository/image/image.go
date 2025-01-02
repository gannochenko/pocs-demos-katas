package image

import (
	"context"

	"backend/internal/database"
	"backend/internal/util/db"
	"backend/internal/util/syserr"

	"gorm.io/gorm"
)

type Repository struct {
	session *gorm.DB
}

func New(session *gorm.DB) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) Create(ctx context.Context, tx *gorm.DB, image *database.Image) error {
	res := db.GetRunner(r.session, tx).WithContext(ctx).Create(image)
	return res.Error
}

func (r *Repository) Update(ctx context.Context, tx *gorm.DB, image *database.Image) error {
	return nil
}

func (r *Repository) List(ctx context.Context, tx *gorm.DB, parameters database.ImageListParameters) ([]database.Image, error) {
	var result []database.Image

	session := db.GetRunner(r.session, tx).WithContext(ctx).Order("created_at desc, id asc")
	session, err := r.applyFilter(session, parameters.Filter)
	if err != nil {
		return nil, syserr.Wrap(err, "could not apply filter")
	}

	err = session.Find(&result).Error
	if err != nil {
		return nil, syserr.Wrap(err, "query failed")
	}

	return result, nil
}

func (r *Repository) Count(ctx context.Context, tx *gorm.DB, parameters database.ImageCountParameters) (int64, error) {
	var result int64

	session := db.GetRunner(r.session, tx).WithContext(ctx).Model(&database.Image{})
	session, err := r.applyFilter(session, parameters.Filter)
	if err != nil {
		return 0, syserr.Wrap(err, "could not apply filter")
	}

	err = session.Count(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *Repository) applyFilter(session *gorm.DB, filter *database.ImageFilter) (*gorm.DB, error) {
	if filter == nil {
		return session, nil
	}

	if filter.CreatedBy != nil {
		session = session.Where("created_by = ?", *filter.CreatedBy)
	}

	return session, nil
}
