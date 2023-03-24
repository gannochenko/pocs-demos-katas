package book

import "levelsgorm/internal/domain/business"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FromBusiness(book *business.Book) (result *Book, err error) {
	return &Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}, nil
}

type GetBooksRequest struct {
	Filter string `json:"filter"`
}

type GetBooksResponse struct {
	Books      []*Book
	Total      int32 `json:"total"`
	PageNumber int32 `json:"page_number"`
}

func GetBooksResponseFromBusiness(response *business.GetBooksResult) (result *GetBooksResponse, err error) {
	var resultBooks []*Book
	for _, book := range response.Books {
		requestBook, err := FromBusiness(book)
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
