package interfaces

import (
	"context"

	databaseBook "be/internal/database/book"
	"gorm.io/gorm"
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
