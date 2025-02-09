package auth

import (
	"encoding/json"
	"io/ioutil"
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
	"api/internal/domain"
	"api/internal/types"
	"api/pkg/ctx"
	"api/pkg/syserr"
)

type Service struct {
	initMu        sync.Mutex
	configService interfaces.ConfigService
	validator     *validator.Validator
	issuer        string
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

	s.issuer = issuerURL.String()

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
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

		_, err = jwtValidator.ValidateToken(r.Context(), token)
		if err != nil {
			if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
				return syserr.NewUnauthorized("missing jwt token")
			}
			if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
				return syserr.NewUnauthorized("jwt token invalid")
			}

			return syserr.Wrap(err, syserr.InternalCode, "could not validate jwt")
		}

		userInfo, err := s.getUserInfo(s.issuer+"userinfo", token)
		if err != nil {
			return syserr.Wrap(err, syserr.InternalCode, "could not retrieve user info")
		}

		r = r.WithContext(ctx.WithUserEmail(r.Context(), userInfo.Email))

		return next(w, r)
	}
}

func (s *Service) getUserInfo(url string, accessToken string) (*domain.UserInfo, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not create a new request")
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not execute the request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not read the response body")
	}

	user := domain.UserInfo{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, syserr.Wrap(err, syserr.InternalCode, "could not unmarshal the response body")
	}

	return &user, nil
}
