package user

import (
	"context"

	"gorm.io/gorm"

	"backend/internal/database"
	"backend/internal/util/db"
	"backend/internal/util/syserr"
)

type Repository struct {
	session *gorm.DB
}

func NewUserRepository(session *gorm.DB) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) List(ctx context.Context, tx *gorm.DB, parameters database.UserListParameters) ([]database.User, error) {
	var result []database.User

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

func (r *Repository) applyFilter(session *gorm.DB, filter *database.UserFilter) (*gorm.DB, error) {
	if filter == nil {
		return session, nil
	}

	if filter.Sup != nil {
		session = session.Where("sup = ?", *filter.Sup)
	}

	return session, nil
}
