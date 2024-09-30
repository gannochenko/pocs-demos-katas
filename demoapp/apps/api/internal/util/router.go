package util

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc AppHandler
}

type Routable interface {
	GetRoutes() map[string]Route
}

func PopulateRouter(router *mux.Router, routables ...Routable) {
	for _, routable := range routables {
		routes := routable.GetRoutes()
		for _, route := range routes {
			router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
				handler := withErrorHandler(withLogger(route.HandlerFunc))
				_ = handler(w, r)
			}).Methods(route.Method)
		}
	}
}
