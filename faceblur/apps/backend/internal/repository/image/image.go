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

func NewImageRepository(session *gorm.DB) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) Create(ctx context.Context, tx *gorm.DB, image *database.Image) error {
	res := db.GetRunner(r.session, tx).WithContext(ctx).Create(image)
	return res.Error
}

func (r *Repository) Update(ctx context.Context, tx *gorm.DB, element *database.ImageUpdate) error {
	if element == nil {
		return syserr.NewBadInput("no element provided")
	}

	updates := make(map[string]interface{})

	if element.URL != nil {
		updates["url"] = element.URL.Value
	}
	if element.OriginalURL != nil {
		updates["original_url"] = element.OriginalURL.Value
	}
	if element.IsProcessed != nil {
		updates["is_processed"] = element.IsProcessed.Value
	}
	if element.IsFailed != nil {
		updates["is_failed"] = element.IsFailed.Value
	}

	if len(updates) == 0 {
		return nil
	}

	result := db.GetRunner(r.session, tx).WithContext(ctx).
		Model(&database.Image{}).
		Where("id = ?", element.ID).
		Updates(updates)

	return result.Error
}

func (r *Repository) List(ctx context.Context, tx *gorm.DB, parameters database.ImageListParameters) ([]database.Image, error) {
	var result []database.Image

	session := db.GetRunner(r.session, tx).WithContext(ctx).Order("uploaded_at desc, created_at desc, id asc")
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

func (r *Repository) Count(ctx context.Context, tx *gorm.DB, parameters database.ImageCountParameters) (int32, error) {
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

	return int32(result), nil
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
