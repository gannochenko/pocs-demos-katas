package book

import (
	"context"
	"loggingerrorhandling/internal/syserr"

	"loggingerrorhandling/internal/domain/business/book"
	databaseBook "loggingerrorhandling/internal/domain/database/book"
	"loggingerrorhandling/internal/logger"
)

type bookRepository interface {
	GetBooks(ctx context.Context, filter string, page int32) (books []*databaseBook.Book, err error)
	GetBookCount(ctx context.Context, filter string) (count int64, err error)
}

type Service struct {
	BookRepository bookRepository
}

func (s *Service) GetBooks(ctx context.Context, filter string, page int32) (result *book.GetBooksResult, err error) {
	logger.Info(ctx, "trying to execute the GetBooks() method")

	result = &book.GetBooksResult{
		PageNumber: page,
		Books:      []*book.Book{},
	}

	panic("oh, no!")

	bookCount, err := s.BookRepository.GetBookCount(ctx, filter)
	if err != nil {
		return nil, syserr.WrapError(err, "could not execute GetBooks()")
	}

	if bookCount == 0 {
		return result, nil
	}

	result.Total = bookCount

	books, err := s.BookRepository.GetBooks(ctx, filter, page)
	if err != nil {
		return nil, err
	}

	if len(books) > 0 {
		for _, dbBook := range books {
			result.Books = append(result.Books, dbBook.ToBusiness())
		}
	}

	return result, nil
}
