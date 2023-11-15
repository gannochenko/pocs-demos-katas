package interfaces

import (
	"context"

	bookDomain "fixtures/internal/domain/book"
)

type BookService interface {
	GetBooks(ctx context.Context, filter string, page int32) (result *bookDomain.GetBooksResult, err error)
}
