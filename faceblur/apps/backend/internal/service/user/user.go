package user

import (
	"context"

	"backend/interfaces"
	"backend/internal/database"
	"backend/internal/domain"
	"backend/internal/util/syserr"
)

type Service struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) GetUserBySUP(ctx context.Context, _ interfaces.SessionHandle, sup string) (*domain.User, error) {
	users, err := s.userRepository.List(ctx, nil, database.UserListParameters{
		Filter: &database.UserFilter{
			Sup: &sup,
		},
	})
	if err != nil {
		return nil, syserr.Wrap(err, "could not get user")
	}

	if len(users) == 0 {
		return nil, syserr.NewNotFound("user was not found", syserr.F("sup", sup))
	}

	domainUser, err := users[0].ToDomain()
	if err != nil {
		return nil, syserr.Wrap(err, "could not convert user to domain")
	}

	return domainUser, nil
}
