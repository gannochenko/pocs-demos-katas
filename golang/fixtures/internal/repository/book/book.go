package book

import (
	"fixtures/constants"
	"fixtures/internal/database/book"
	"gorm.io/gorm"
)

const (
	TableName = "books"
)

type Repository struct {
	Session *gorm.DB
}

func (r *Repository) List(parameters *book.ListParameters) (books []*book.Book, err error) {
	runner := r.Session.Table(TableName)

	r.applyFilter(runner, parameters)

	runner.Offset(int(parameters.Page * constants.PageNavPageSize)).Limit(constants.PageNavPageSize).Find(&books)
	return books, nil
}

func (r *Repository) Count(parameters *book.ListParameters) (count int64, err error) {
	runner := r.Session.Table(TableName)

	r.applyFilter(runner, parameters)

	runner.Select("id").Count(&count)
	return count, nil
}

func (r *Repository) applyFilter(runner *gorm.DB, parameters *book.ListParameters) {
	if parameters.Filter == nil {
		return
	}

	filter := parameters.Filter

	if filter.Title != nil {
		runner.Where("title like ?", *filter.Title)
	}

	if filter.AuthorID != nil {
		runner.Where("author_id = ?", *filter.AuthorID)
	}
}
