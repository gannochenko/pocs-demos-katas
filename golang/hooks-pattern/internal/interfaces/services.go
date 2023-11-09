package interfaces

import (
	"context"

	bookDomain "hookspattern/internal/domain/book"
)

type BookService interface {
	GetBooks(ctx context.Context, filter string, page int32) (result *bookDomain.GetBooksResult, err error)
	DeleteBook(ctx context.Context, bookID string) (err error)
}

type HooksService interface {
	On(eventName string, handler func(ctx context.Context, args interface{}))
	Trigger(ctx context.Context, eventName string, args interface{})
}
