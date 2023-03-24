package book

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"levelsgorm/internal/rest/book"
)

type Controller struct{}

func (c *Controller) GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	jsonBody := book.Request{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	io.WriteString(responseWriter, "This is my website!\n")
}
