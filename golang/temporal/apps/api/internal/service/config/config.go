package config

import (
	"sync"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"api/internal/domain"
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

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("/etc/config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	var config domain.Config

	if err := defaults.Set(&config); err != nil {
		return err
	}

	if err := v.Unmarshal(&config); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		return err
	}

	s.Config = &config

	return nil
}
