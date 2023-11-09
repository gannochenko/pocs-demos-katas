package book

import (
	"hookspattern/internal/domain/book"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	AuthorID string `json:"author_id"`
}

func FromDomain(book *book.Book) (result *Book, err error) {
	return &Book{
		ID:       book.ID,
		Title:    book.Title,
		AuthorID: book.AuthorID,
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

func FromDomainGetBooksResponse(response *book.GetBooksResult) (result *GetBooksResponse, err error) {
	var resultBooks []*Book
	for _, businessBook := range response.Books {
		requestBook, err := FromDomain(businessBook)
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

type DeleteBookRequest struct {
	BookID string `json:"book_id"`
}

type DeleteBookResponse struct {
}

func FromDomainDeleteBookResponse(response *book.DeleteBookResult) (result *DeleteBookResponse, err error) {
	return &DeleteBookResponse{}, nil
}
