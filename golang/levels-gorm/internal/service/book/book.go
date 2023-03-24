package book

import (
	"levelsgorm/internal/domain/business"
)

type Service struct{}

func (s *Service) GetBooks(filter string) (result *business.GetBooksResult, err error) {
	result = &business.GetBooksResult{
		Books: []*business.Book{
			{
				ID:        "123",
				Title:     "1984",
				Author:    "Oruell",
				IssueYear: 1949,
			},
		},
		Total:      100,
		PageNumber: 1,
	}

	return result, nil
}
