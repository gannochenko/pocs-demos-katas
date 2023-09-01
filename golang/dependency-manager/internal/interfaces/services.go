package interfaces

import "dependencymanager/internal/domain/business/book"

type BookService interface {
	GetBooks(filter string, page int32) (result *book.GetBooksResult, err error)
}
