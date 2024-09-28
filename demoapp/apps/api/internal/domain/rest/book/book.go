package book

import (
	"api/internal/domain/business/book"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FromBusiness(book *book.Book) (result *Book, err error) {
	return &Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}, nil
}

type GetBooksRequest struct {
	Filter string `json:"filter"`
	Page   int32  `json:"page"`
}

type GetBooksResponse struct {
	Books      []*Book `json:"books"`
	Total      int64   `json:"total"`
	PageNumber int32   `json:"page_number"`
}

func FromBusinessGetBooksResponse(response *book.GetBooksResult) (result *GetBooksResponse, err error) {
	resultBooks := []*Book{}
	for _, businessBook := range response.Books {
		requestBook, err := FromBusiness(businessBook)
		if err != nil {
			return nil, err
		}
		resultBooks = append(resultBooks, requestBook)
	}

	return &GetBooksResponse{
		Books:      resultBooks,
		Total:      response.Total,
		PageNumber: response.PageNumber,
	}, nil
}
