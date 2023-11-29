package book

import (
	"context"

	"be/internal/constants"
	"be/internal/domain/book"
	"be/internal/interfaces"
)

type Service struct {
	BookRepository interfaces.BookRepository
	HooksService   interfaces.HooksService
}

func (s *Service) GetBooks(ctx context.Context, filter string, page int32) (result *book.GetBooksResult, err error) {
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
			result.Books = append(result.Books, dbBook.ToDomain())
		}
	}

	return result, nil
}

func (s *Service) DeleteBook(ctx context.Context, bookID string) error {
	err := s.BookRepository.DeleteBook(bookID)
	if err != nil {
		return err
	}

	s.HooksService.Trigger(ctx, constants.EventOnAfterBookDelete, []string{bookID})

	return nil
}
