package interfaces

import (
	"context"

	"gorm.io/gorm"
	databaseBook "hookspattern/internal/database/book"
)

type BookRepository interface {
	GetBooks(filter string, page int32) (books []*databaseBook.Book, err error)
	GetBookCount(filter string) (count int64, err error)
	DeleteBook(bookID string) (err error)
	GetAuthorIDsByBookIDsSQL(bookIDs []string) *gorm.DB
}

type AuthorRepository interface {
	RefreshHasBooksFlag(ctx context.Context, condition interface{}) (err error)
}
