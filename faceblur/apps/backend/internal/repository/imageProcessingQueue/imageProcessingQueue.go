package imageProcessingQueue

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"gorm.io/gorm"

	"backend/internal/database"
	"backend/internal/util/db"
	"backend/internal/util/syserr"
)

type Repository struct {
	session *gorm.DB
}

func NewImageProcessingQueueRepository(session *gorm.DB) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) Create(ctx context.Context, tx *gorm.DB, element *database.ImageProcessingQueue) error {
	res := db.GetRunner(r.session, tx).WithContext(ctx).Create(element)
	return res.Error
}

func (r *Repository) Update(ctx context.Context, tx *gorm.DB, element *database.ImageProcessingQueue) error {
	return nil
}

func (r *Repository) List(ctx context.Context, tx *gorm.DB, parameters database.ImageProcessingQueueListParameters) ([]database.ImageProcessingQueue, error) {
	var result []database.ImageProcessingQueue

	session := db.GetRunner(r.session, tx).WithContext(ctx).Order("created_at desc")
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

func (r *Repository) applyFilter(session *gorm.DB, filter *database.ImageProcessingQueueFilter) (*gorm.DB, error) {
	if filter == nil {
		return session, nil
	}

	if filter.CreatedBy != nil {
		session = session.Where("created_by = ?", *filter.CreatedBy)
	}

	if filter.IsFailed != nil {
		session = session.Where("is_failed = ?", *filter.IsFailed)
	}

	if filter.IsCompleted != nil {
		session = session.Where(fmt.Sprintf("completed_at is %s null", lo.Ternary(*filter.IsCompleted, "not", "")))
	}

	return session, nil
}
