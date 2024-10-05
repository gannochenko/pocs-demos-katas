package util

import (
	"net/http"

	"github.com/gorilla/mux"

	"api/interfaces"
	"api/internal/domain"
	"api/internal/types"
	httpUtil "api/internal/util/http"
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

func PopulateRouter(router *mux.Router, config *domain.Config, authService interfaces.AuthService, routables ...Routable) {
	for _, routable := range routables {
		routes := routable.GetRoutes()
		for _, route := range routes {
			router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
				handler := route.HandlerFunc
				if route.Protected {
					handler = authService.WithAuth(handler)
				}
				handler = httpUtil.WithCORS(config, WithErrorHandler(WithLogger(handler)))
				_ = handler(w, r)
			}).Methods(route.Method, http.MethodOptions)
		}
	}
}
