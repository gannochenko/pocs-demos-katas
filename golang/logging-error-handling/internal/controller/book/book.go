package book

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	bookBusiness "loggingerrorhandling/internal/domain/business/book"
	"loggingerrorhandling/internal/domain/rest/book"
)

type bookService interface {
	GetBooks(filter string, page int32) (result *bookBusiness.GetBooksResult, err error)
}

type Controller struct {
	BookService bookService
}

func (c *Controller) GetBooks(_ http.ResponseWriter, request *http.Request) ([]byte, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	jsonBody := book.GetBooksRequest{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return []byte{}, err
	}

	result, err := c.BookService.GetBooks(jsonBody.Filter, jsonBody.Page)
	if err != nil {
		return []byte{}, err
	}

	bookResponse, err := book.FromBusinessGetBooksResponse(result)
	if err != nil {
		return []byte{}, err
	}

	responseBody, err := json.Marshal(bookResponse)
	if err != nil {
		return []byte{}, err
	}

	return []byte{}, fmt.Errorf("Oh My Glob")

	return responseBody, nil
}
