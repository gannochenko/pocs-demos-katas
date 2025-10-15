package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"

	"gateway/internal/domain"
)

type Service struct {
	initMu sync.Mutex
	Config *domain.Config
}

func NewConfigService() *Service {
	return &Service{}
}

func (s *Service) GetConfig() *domain.Config {
	return s.Config
}

func (s *Service) LoadConfig() error {
	if s.Config != nil {
		return nil
	}

	s.initMu.Lock()
	defer s.initMu.Unlock()

	var config domain.Config
	err := envconfig.Process("GATEWAY", &config)
	if err != nil {
		return err
	}

	s.Config = &config

	return nil
}
