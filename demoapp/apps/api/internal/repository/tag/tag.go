package tag

import (
	"context"

	"gorm.io/gorm"

	"api/internal/dto"
	"api/internal/util/db"
)

const (
	TableName = "tags"
)

type Repository struct {
	session *gorm.DB
}

func New(session *gorm.DB) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) ListTags(ctx context.Context, tx *gorm.DB, parameters *dto.ListTagsParameters) (result []*dto.Tag, err error) {
	runner := db.GetRunner(r.session, tx).WithContext(ctx).Table(TableName)
	runner = r.applyFilter(runner, parameters)

	if parameters != nil {
		runner = db.ApplyPagination(runner, parameters.Pagination)
	}

	queryResult := runner.Find(&result)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	return result, nil
}

func (r *Repository) CountTags(ctx context.Context, tx *gorm.DB, parameters *dto.ListTagsParameters) (count int64, err error) {
	runner := db.GetRunner(r.session, tx).WithContext(ctx).Table(TableName)
	runner = r.applyFilter(runner, parameters)

	queryResult := runner.Model(&dto.Pet{}).Count(&count)
	if queryResult.Error != nil {
		return 0, queryResult.Error
	}
	return count, nil
}

func (r *Repository) applyFilter(runner *gorm.DB, parameters *dto.ListTagsParameters) *gorm.DB {
	if parameters != nil {
		if parameters.Filter != nil {
			//filter := parameters.Filter
		}
	}

	return runner
}