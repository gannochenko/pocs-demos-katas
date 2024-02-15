package http

import (
	"net/http"
)

func RequestWriter(controllerFn func(w http.ResponseWriter, r *http.Request) ([]byte, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseBody, err := controllerFn(w, r)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if len(responseBody) > 0 {
			_, _ = w.Write(responseBody)
		}
	})
}
