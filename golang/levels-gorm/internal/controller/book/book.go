package book

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"levelsgorm/internal/domain/business"
	"levelsgorm/internal/rest/book"
)

type bookService interface {
	GetBooks(filter string) (result *business.GetBooksResult, err error)
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

	result, err := c.BookService.GetBooks(jsonBody.Filter)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookResponse, err := book.GetBooksResponseFromBusiness(result)
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
