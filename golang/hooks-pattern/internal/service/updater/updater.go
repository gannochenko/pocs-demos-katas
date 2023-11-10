package updater

import (
	"context"
	"fmt"
	"net/http"

	"hookspattern/internal/constants"
	"hookspattern/internal/interfaces"
	"hookspattern/pkg/slice"
)

type CtxKeyType string

const (
	CtxKey CtxKeyType = "updater-values"
)

type Values struct {
	affectedBooks []string
}

type Service struct {
	hooksService     interfaces.HooksService
	authorRepository interfaces.AuthorRepository
}

func New(hooksService interfaces.HooksService, authorRepository interfaces.AuthorRepository) *Service {
	return &Service{
		hooksService:     hooksService,
		authorRepository: authorRepository,
	}
}

func (s *Service) WithValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, CtxKey, &Values{})
}

func (s *Service) GetValues(ctx context.Context) (*Values, error) {
	if value, ok := ctx.Value(CtxKey).(*Values); ok {
		return value, nil
	}

	return nil, fmt.Errorf("values missing from the context")
}

func (s *Service) Init() {
	s.hooksService.On(constants.EventOnAfterBookDelete, func(ctx context.Context, args interface{}) {
		values, err := s.GetValues(ctx)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}

		if ids, ok := args.([]string); ok {
			values.affectedBooks = slice.Merge(values.affectedBooks, ids)
		} else {
			fmt.Println("the argument is not of correct type, this is a noop")
			return
		}

		fmt.Println("Deletion event processed!")
	})
}

func (s *Service) Process(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context is done, this is a noop")
		return
	default:
	}

	fmt.Println("Processing!")
	values, err := s.GetValues(ctx)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	defer func() {
		values.affectedBooks = []string{}
	}()

	fmt.Printf("%v", values)
}

func (s *Service) GetHTTPMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := s.WithValue(r.Context())
			// fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)

			next.ServeHTTP(w, r.WithContext(ctx))

			s.Process(ctx)
		})
	}
}
