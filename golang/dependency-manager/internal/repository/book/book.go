package book

import (
	"dependencymanager/internal/domain/database/book"
	"gorm.io/gorm"
)

const (
	TableName = "books"
	PageSize  = 5
)

type Repository struct {
	Session *gorm.DB
}

func New(session *gorm.DB) *Repository {
	return &Repository{
		Session: session,
	}
}

func (r *Repository) GetBooks(filter string, page int32) (books []*book.Book, err error) {
	runner := r.Session.Table(TableName)
	if filter != "" {
		runner = runner.Where("title like ?", filter)
	}
	runner.Offset(int(page * PageSize)).Limit(PageSize).Find(&books)
	return books, nil
}

func (r *Repository) GetBookCount(filter string) (count int64, err error) {
	runner := r.Session.Table(TableName)
	if filter != "" {
		runner = runner.Where("title like ?", filter)
	}
	runner.Select("id").Count(&count)
	return count, nil
}
