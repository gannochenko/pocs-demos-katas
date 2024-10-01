package util

import (
	"net/http"

	"github.com/gorilla/mux"

	"api/interfaces"
	"api/internal/types"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc types.Handler
	Protected   bool
}

type Routable interface {
	GetRoutes() map[string]Route
}

func PopulateRouter(router *mux.Router, authService interfaces.AuthService, routables ...Routable) {
	for _, routable := range routables {
		routes := routable.GetRoutes()
		for _, route := range routes {
			router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
				handler := route.HandlerFunc
				//if route.Protected {
				//	handler = authService.WithAuth(handler)
				//}
				handler = withErrorHandler(withLogger(handler))
				_ = handler(w, r)
			}).Methods(route.Method)
		}
	}
}
