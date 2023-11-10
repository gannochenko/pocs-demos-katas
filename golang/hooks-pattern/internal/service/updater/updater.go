package updater

import (
	"context"
	"fmt"

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
	})
}

func (s *Service) Process() {

}

func (s *Service) GetHTTPMiddleware() {

}
