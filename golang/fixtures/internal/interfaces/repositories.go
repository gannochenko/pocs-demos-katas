package interfaces

import (
	"fixtures/internal/database/author"
	"fixtures/internal/database/book"
)

type BookRepository interface {
	List(parameters *book.ListParameters) (books []*book.Book, err error)
	Count(parameters *book.ListParameters) (count int64, err error)
}

type AuthorRepository interface {
	List(parameters *author.ListParameters) (authors []*author.Author, err error)
}
