package main

import (
	"syscall/js"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/norunners/vert"
)

// https://stackoverflow.com/questions/68798120/imports-syscall-js-build-constraints-exclude-all-go-files-in-usr-local-go-src
// https://reintech.io/blog/a-guide-to-gos-syscall-js-package-go-and-webassembly
// https://github.com/norunners/vert

type GetItemsRequest struct {
	Amount int32
}

type Item struct {
	ID    int32  `js:"id"`
	Title string `js:"title"`
	Date  string `js:"date"`
}

type GetItemsResponse struct {
	Error string `js:"error"`
	Items []Item `js:"items"`
}

func convertGetItemsRequest(jsRequest js.Value) *GetItemsRequest {
	return &GetItemsRequest{
		Amount: int32(jsRequest.Get("amount").Int()),
	}
}

func getItems(request *GetItemsRequest) *GetItemsResponse {
	response := &GetItemsResponse{}

	if request.Amount <= 0 {
		response.Error = "request amount should be a positive number"
		return response
	}

	items := make([]Item, 0)

	for i := int32(0); i < request.Amount; i++ {
		items = append(items, Item{
			ID:    gofakeit.Int32(),
			Title: gofakeit.BeerName(),
			Date:  gofakeit.Date().Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	response.Items = items

	return response
}

func main() {
	js.Global().Set("getItems", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return vert.ValueOf(getItems(convertGetItemsRequest(args[0]))).JSValue()
	}))

	// block forever
	select {}
}
