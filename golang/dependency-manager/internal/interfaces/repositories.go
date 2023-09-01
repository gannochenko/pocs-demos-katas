package interfaces

import (
	"dependencymanager/internal/domain/database/book"
)

type BookRepository interface {
	GetBooks(filter string, page int32) (books []*book.Book, err error)
	GetBookCount(filter string) (count int64, err error)
}
