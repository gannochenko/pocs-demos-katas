package updater

import (
	"context"
	"fmt"

	"hookspattern/internal/constants"
	"hookspattern/internal/interfaces"
)

type Service struct {
	hooksService interfaces.HooksService
	// todo: +authors repository here
}

func New(hooksService interfaces.HooksService) *Service {
	return &Service{
		hooksService: hooksService,
	}
}

func (s *Service) Init() {
	s.hooksService.On(constants.EventOnAfterBookDelete, func(ctx context.Context, args interface{}) {
		fmt.Println("Called:")
		fmt.Printf("%v\n", args)
	})
}

// todo: + middleware maker
