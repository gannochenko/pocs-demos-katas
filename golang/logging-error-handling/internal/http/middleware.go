package http

import (
	"encoding/json"
	"net/http"
)

func RequestWriter(controllerFn func(w http.ResponseWriter, r *http.Request) ([]byte, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody, err := controllerFn(w, r)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			errorResponse := &ErrorResponse{
				Error: err.Error(),
			}
			responseBody, _ := json.Marshal(errorResponse)
			_, _ = w.Write(responseBody)

			return
		} else {
			w.WriteHeader(http.StatusOK)
			if len(responseBody) > 0 {
				_, _ = w.Write(responseBody)
			}
		}
	})
}
