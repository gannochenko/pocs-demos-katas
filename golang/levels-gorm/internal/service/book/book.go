package book

import (
	"levelsgorm/internal/domain/business/book"
)

type Service struct{}

func (s *Service) GetBooks(filter string) (result *book.GetBooksResult, err error) {
	result = &book.GetBooksResult{
		Books: []*book.Book{
			{
				ID:        "123",
				Title:     "1984",
				Author:    "Orwell",
				IssueYear: 1949,
			},
		},
		Total:      100,
		PageNumber: 1,
	}

	return result, nil
}
