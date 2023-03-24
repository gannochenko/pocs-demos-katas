package book

import (
	"io"
	"net/http"
)

type Controller struct{}

func (c *Controller) GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	//ctx := request.Context()

	io.WriteString(responseWriter, "This is my website!\n")
}
