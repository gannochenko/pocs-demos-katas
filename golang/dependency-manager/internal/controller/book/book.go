package book

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	bookBusiness "dependencymanager/internal/domain/business/book"
	"dependencymanager/internal/domain/rest/book"
)

type bookService interface {
	GetBooks(filter string, page int32) (result *bookBusiness.GetBooksResult, err error)
}

type Controller struct {
	BookService bookService
}

func (c *Controller) GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	jsonBody := book.GetBooksRequest{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.BookService.GetBooks(jsonBody.Filter, jsonBody.Page)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookResponse, err := book.FromBusinessGetBooksResponse(result)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(bookResponse)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	_, err = responseWriter.Write(responseBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
}
