package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"

	"backend/interfaces"
	"backend/internal/domain"
	"backend/internal/util/syserr"
)

type CustomClaims struct {
	Permissions []string `json:"permissions"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

type Service struct {
	initMu        sync.Mutex
	configService interfaces.ConfigService
	loggerService interfaces.LoggerService
	validator     *validator.Validator
	issuerURL     string
}

func NewAuthService(configService interfaces.ConfigService, loggerService interfaces.LoggerService) *Service {
	return &Service{
		configService: configService,
		loggerService: loggerService,
	}
}

func (s *Service) ExtractToken(ctx context.Context) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	tokens := md["authorization"]

	if len(tokens) == 0 {
		return "", syserr.NewBadInput("missing token")
	}

	authHeaderParts := strings.Fields(tokens[0])
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", syserr.NewBadInput("authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

func (s *Service) ValidateToken(ctx context.Context, token string) (string, int64, error) {
	jwtValidator, err := s.getValidator()
	if err != nil {
		return "", 0, syserr.Wrap(err, "could not get validator")
	}

	claims, err := jwtValidator.ValidateToken(ctx, token)
	if err != nil {
		if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
			return "", 0, syserr.NewUnauthorized("missing jwt token")
		}
		if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
			return "", 0, syserr.NewUnauthorized("jwt token invalid")
		}

		return "", 0, err
	}

	validatedClaims := claims.(*validator.ValidatedClaims)

	return validatedClaims.RegisteredClaims.Subject, validatedClaims.RegisteredClaims.Expiry, nil
}

func (s *Service) getValidator() (*validator.Validator, error) {
	if s.validator != nil {
		return s.validator, nil
	}

	s.initMu.Lock()
	defer s.initMu.Unlock()

	config, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not load config")
	}

	issuerURL, err := url.Parse("https://" + config.Auth0.Domain + "/")
	if err != nil {
		return nil, syserr.Wrap(err, "could not create issuer URL")
	}

	s.issuerURL = issuerURL.String()

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{config.Auth0.Audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return new(CustomClaims)
		}),
	)
	if err != nil {
		return nil, syserr.Wrap(err, "could not setup jwt validator")
	}

	s.validator = jwtValidator

	return jwtValidator, nil
}

func (s *Service) GetUserInfo(ctx context.Context, accessToken string) (*domain.RemoteUserInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%suserinfo", s.issuerURL), nil)
	if err != nil {
		return nil, syserr.Wrap(err, "could not create a new request")
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, syserr.Wrap(err, "could not execute the request")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			s.loggerService.LogError(ctx, syserr.Wrap(err, "could not close response body reader"))
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, syserr.Wrap(err, "could not read the response body")
	}

	user := domain.RemoteUserInfo{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, syserr.Wrap(err, "could not unmarshal the response body")
	}

	return &user, nil
}
