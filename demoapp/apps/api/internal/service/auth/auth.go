package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"context"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/pkg/errors"

	"api/interfaces"
	"api/internal/types"
	"api/pkg/syserr"
)

type Service struct {
	initMu        sync.Mutex
	configService interfaces.ConfigService
	validator     *validator.Validator
}

func New(configService interfaces.ConfigService) *Service {
	return &Service{
		configService: configService,
	}
}

type CustomClaims struct {
	Permissions []string `json:"permissions"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

//func (c CustomClaims) HasPermissions(expectedClaims []string) bool {
//	if len(expectedClaims) == 0 {
//		return false
//	}
//	for _, scope := range expectedClaims {
//		if !helpers.Contains(c.Permissions, scope) {
//			return false
//		}
//	}
//	return true
//}

//func ValidatePermissions(expectedClaims []string, next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
//		claims := token.CustomClaims.(*CustomClaims)
//		if !claims.HasPermissions(expectedClaims) {
//			errorMessage := ErrorMessage{Message: permissionDeniedErrorMessage}
//			if err := helpers.WriteJSON(w, http.StatusForbidden, errorMessage); err != nil {
//				log.Printf("Failed to write error message: %v", err)
//			}
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}

func (s *Service) GetValidator() (*validator.Validator, error) {
	if s.validator != nil {
		return s.validator, nil
	}

	s.initMu.Lock()
	defer s.initMu.Unlock()

	conf, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not extract configService")
	}

	domain := conf.Auth0.Domain
	audience := conf.Auth0.Audience

	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.PS256,
		issuerURL.String(),
		[]string{audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return new(CustomClaims)
		}),
	)
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not setup jwt validator")
	}

	s.validator = jwtValidator

	return jwtValidator, nil
}

func (s *Service) WithAuth(next types.Handler) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		jwtValidator, err := s.GetValidator()
		if err != nil {
			return err
		}

		token, err := jwtmiddleware.AuthHeaderTokenExtractor(r)
		if err != nil {
			return syserr.Wrap(err, syserr.BadInputCode, "could not extract jwt token")
		}

		claims, err := jwtValidator.ValidateToken(r.Context(), token)
		if err != nil {
			if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
				return syserr.NewUnauthorized("missing jwt token")
			}
			if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
				return syserr.NewUnauthorized("jwt token invalid")
			}

			return syserr.Wrap(err, syserr.InternalCode, "could not validate jwt")
		}

		fmt.Printf("%v", claims)

		return next(w, r)
	}
}
