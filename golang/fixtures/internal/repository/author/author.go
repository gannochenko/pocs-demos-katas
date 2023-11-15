package author

import (
	"fixtures/constants"
	"fixtures/internal/database/author"
	"gorm.io/gorm"
)

const (
	TableName = "authors"
)

type Repository struct {
	Session *gorm.DB
}

func (r *Repository) List(parameters *author.ListParameters) (authors []*author.Author, err error) {
	runner := r.Session.Table(TableName)

	r.applyFilter(runner, parameters)

	runner.Offset(int(parameters.Page * constants.PageNavPageSize)).Limit(constants.PageNavPageSize).Find(&authors)
	return authors, nil
}

func (r *Repository) applyFilter(runner *gorm.DB, parameters *author.ListParameters) {
	if parameters.Filter == nil {
		return
	}

	filter := parameters.Filter

	if filter.ID != nil {
		runner.Where("id = ?", *filter.ID)
	}
}
