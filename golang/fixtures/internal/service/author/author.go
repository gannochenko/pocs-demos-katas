package author

import (
	"context"

	"fixtures/internal/database/author"
	"fixtures/internal/database/book"
	authorDomain "fixtures/internal/domain/author"
	"fixtures/internal/interfaces"
)

type Service struct {
	bookRepository   interfaces.BookRepository
	authorRepository interfaces.AuthorRepository
}

func New(bookRepository interfaces.BookRepository, authorRepository interfaces.AuthorRepository) *Service {
	return &Service{
		bookRepository:   bookRepository,
		authorRepository: authorRepository,
	}
}

func (s *Service) GetBooksAndAuthor(ctx context.Context, authorID string) (result *authorDomain.GetBooksAndAuthorResult, err error) {
	result = &authorDomain.GetBooksAndAuthorResult{}

	books, err := s.bookRepository.List(&book.ListParameters{
		Filter: &book.ListParametersFilter{
			AuthorID: &authorID,
		},
	})
	if err != nil {
		return nil, err
	}

	authors, err := s.authorRepository.List(&author.ListParameters{
		Filter: &author.ListParametersFilter{
			ID: &authorID,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(authors) > 0 {
		result.Author = authors[0].ToDomain()
	}

	if len(books) > 0 {
		for _, dbBook := range books {
			result.Books = append(result.Books, dbBook.ToDomain())
		}
	}

	return result, nil
}
