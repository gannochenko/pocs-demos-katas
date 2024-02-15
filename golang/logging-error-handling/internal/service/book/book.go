package book

import (
	"loggingerrorhandling/internal/domain/business/book"
	databaseBook "loggingerrorhandling/internal/domain/database/book"
)

type bookRepository interface {
	GetBooks(filter string, page int32) (books []*databaseBook.Book, err error)
	GetBookCount(filter string) (count int64, err error)
}

type Service struct {
	BookRepository bookRepository
}

func (s *Service) GetBooks(filter string, page int32) (result *book.GetBooksResult, err error) {
	result = &book.GetBooksResult{
		PageNumber: page,
		Books:      []*book.Book{},
	}

	bookCount, err := s.BookRepository.GetBookCount(filter)
	if err != nil {
		return nil, err
	}

	if bookCount == 0 {
		return result, nil
	}

	result.Total = bookCount

	books, err := s.BookRepository.GetBooks(filter, page)
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
