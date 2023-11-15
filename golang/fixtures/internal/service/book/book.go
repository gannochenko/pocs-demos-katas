package book

import (
	"context"

	databaseBook "fixtures/internal/database/book"
	"fixtures/internal/domain/book"
	"fixtures/internal/interfaces"
)

type Service struct {
	BookRepository interfaces.BookRepository
}

func (s *Service) GetBooks(ctx context.Context, filter string, page int32) (result *book.GetBooksResult, err error) {
	result = &book.GetBooksResult{
		PageNumber: page,
		Books:      []*book.Book{},
	}

	bookCount, err := s.BookRepository.Count(&databaseBook.ListParameters{
		Page: page,
		Filter: &databaseBook.ListParametersFilter{
			Title: &filter,
		},
	})
	if err != nil {
		return nil, err
	}

	if bookCount == 0 {
		return result, nil
	}

	result.Total = bookCount

	books, err := s.BookRepository.List(&databaseBook.ListParameters{
		Page: page,
		Filter: &databaseBook.ListParametersFilter{
			Title: &filter,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(books) > 0 {
		for _, dbBook := range books {
			result.Books = append(result.Books, dbBook.ToDomain())
		}
	}

	return result, nil
}
