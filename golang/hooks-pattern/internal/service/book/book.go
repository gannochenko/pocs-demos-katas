package book

import (
	"hookspattern/internal/domain/book"
	"hookspattern/internal/interfaces"
)

type Service struct {
	BookRepository interfaces.BookRepository
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
			result.Books = append(result.Books, dbBook.ToDomain())
		}
	}

	return result, nil
}

func (s *Service) DeleteBook(bookID string) error {
	return s.BookRepository.DeleteBook(bookID)
}
