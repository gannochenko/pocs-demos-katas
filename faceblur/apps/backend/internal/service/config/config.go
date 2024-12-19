package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"

	"backend/internal/domain"
)

type Config struct {
	initMu sync.Mutex
	config *domain.Config
}

func NewConfigService() *Config {
	return &Config{}
}

func (c *Config) GetConfig() (*domain.Config, error) {
	if c.config != nil {
		return c.config, nil
	}

	c.initMu.Lock()
	defer c.initMu.Unlock()

	var config domain.Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	c.config = &config

	return c.config, nil
}
