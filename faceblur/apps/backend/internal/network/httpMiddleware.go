package network

import (
	"net/http"

	"github.com/google/uuid"

	"backend/internal/domain"
	ctxUtil "backend/internal/util/ctx"
	"backend/internal/util/types"
)

func withHTTPLogger(next types.HTTPHandler) types.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return next(w, r)
	}
}

func withHTTPErrorHandler(next types.HTTPHandler) types.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return next(w, r)
	}
}

func withHTTPContext(next types.HTTPHandler) types.HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return next(w, r.WithContext(ctxUtil.WithOperationID(r.Context(), uuid.NewString())))
	}
}

// withCorsMiddleware adds CORS headers. Normally Kubernetes ingress or CDN takes care of that, but for the dev purposes we add it here as well.
func withCorsMiddleware(next http.Handler, config *domain.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := config.HTTP.Cors.Origin[0]

		// Handle CORS headers
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			// Handle preflight OPTIONS request
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
