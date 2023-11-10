package interfaces

import (
	"context"

	databaseBook "hookspattern/internal/database/book"
)

type BookRepository interface {
	GetBooks(filter string, page int32) (books []*databaseBook.Book, err error)
	GetBookCount(filter string) (count int64, err error)
	DeleteBook(bookID string) (err error)
}

type AuthorRepository interface {
	RefreshHasBooksFlag(ctx context.Context, condition interface{}) (err error)
}
