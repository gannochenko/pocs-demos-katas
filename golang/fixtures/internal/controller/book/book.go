package book

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fixtures/internal/interfaces"
	"fixtures/internal/rest/book"
)

type Controller struct {
	BookService interfaces.BookService
}

func (c *Controller) GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

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

	result, err := c.BookService.GetBooks(ctx, jsonBody.Filter, jsonBody.Page)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	bookResponse, err := book.FromDomainGetBooksResponse(result)
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
