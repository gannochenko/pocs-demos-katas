package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"

	"api/internal/domain"
)

type Config struct {
	mu     sync.Mutex
	config *domain.Config
}

func NewConfigService() *Config {
	return &Config{}
}

func (c *Config) GetConfig() (*domain.Config, error) {
	if c.config != nil {
		return c.config, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	var config domain.Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	c.config = &config

	return c.config, nil
}
