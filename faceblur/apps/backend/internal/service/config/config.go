package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"

	"backend/internal/domain"
)

type Service struct {
	initMu sync.Mutex
	config *domain.Config
}

func NewConfigService() *Service {
	return &Service{}
}

func (s *Service) GetConfig() (*domain.Config, error) {
	if s.config != nil {
		return s.config, nil
	}

	s.initMu.Lock()
	defer s.initMu.Unlock()

	var config domain.Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	s.config = &config

	return s.config, nil
}
