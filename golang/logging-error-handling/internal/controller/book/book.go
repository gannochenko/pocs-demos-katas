package book

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	bookBusiness "loggingerrorhandling/internal/domain/business/book"
	"loggingerrorhandling/internal/domain/rest/book"
)

type bookService interface {
	GetBooks(ctx context.Context, filter string, page int32) (result *bookBusiness.GetBooksResult, err error)
}

type Controller struct {
	BookService bookService
}

func (c *Controller) GetBooks(_ http.ResponseWriter, request *http.Request) (body []byte, err error) {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	jsonBody := book.GetBooksRequest{}
	err = json.Unmarshal(requestBody, &jsonBody)
	if err != nil {
		return []byte{}, err
	}

	result, err := c.BookService.GetBooks(request.Context(), jsonBody.Filter, jsonBody.Page)
	if err != nil {
		return []byte{}, err
	}

	bookResponse, err := book.FromBusinessGetBooksResponse(result)
	if err != nil {
		return []byte{}, err
	}

	body, err = json.Marshal(bookResponse)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
