package util

import (
	"net/http"

	openapi "api/internal/bullshit"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func withErrorHandler(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			_ = openapi.EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
		}
	}
}

//func DefaultErrorHandler(w http.ResponseWriter, _ *http.Request, err error, result *ImplResponse) {
//	var parsingErr *ParsingError
//	if ok := errors.As(err, &parsingErr); ok {
//		// Handle parsing errors
//		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
//		return
//	}
//
//	var requiredErr *RequiredError
//	if ok := errors.As(err, &requiredErr); ok {
//		// Handle missing required errors
//		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusUnprocessableEntity), w)
//		return
//	}
//
//	// Handle all other errors
//	_ = EncodeJSONResponse(err.Error(), &result.Code, w)
//}
