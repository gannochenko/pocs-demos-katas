package book

import (
	"dependencymanager/internal/domain/business/book"
	"dependencymanager/internal/interfaces"
	"gorm.io/gorm"
)

type Service struct {
	BookRepository interfaces.BookRepository
	session        *gorm.DB // the session is needed here to manage transactions, as we dont have a separate service for it
}

func New(session *gorm.DB, bookRepository interfaces.BookRepository) *Service {
	return &Service{
		BookRepository: bookRepository,
		session:        session,
	}
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
