package interfaces

import bookDomain "hookspattern/internal/domain/book"

type BookService interface {
	GetBooks(filter string, page int32) (result *bookDomain.GetBooksResult, err error)
	DeleteBook(bookID string) (err error)
}
