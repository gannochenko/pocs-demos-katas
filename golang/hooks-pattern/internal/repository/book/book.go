package book

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hookspattern/internal/database/book"
)

const (
	TableName = "books"
	PageSize  = 5
)

type Repository struct {
	Session *gorm.DB
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

func (r *Repository) DeleteBook(bookID string) (err error) {
	id, err := uuid.Parse(bookID)
	if err != nil {
		return err
	}

	r.Session.Delete(&book.Book{
		ID: id,
	})

	return nil
}

func (r *Repository) GetAuthorIDsByBookIDsSQL(bookIDs []string) *gorm.DB {
	return r.Session.Model(&book.Book{}).Select([]string{"author_id"}).Where("id IN ? AND deleted_at IS NULL", bookIDs)
}
