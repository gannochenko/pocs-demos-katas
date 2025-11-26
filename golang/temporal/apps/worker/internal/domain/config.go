package domain

import (
	"crypto/tls"

	"go.temporal.io/sdk/client"
)

type HTTPConfig struct {
	Addr string `mapstructure:"addr" default:"localhost:8080" validate:"required,hostname_port"`
}

type ServiceConfig struct {
	Name    string `mapstructure:"name" default:"api" validate:"required"`
	Version string `mapstructure:"version" default:"1.0.0" validate:"required"`
}

type TemporalConfig struct {
	Addr string `mapstructure:"addr" default:"localhost:7233" validate:"required,hostname_port"`
	TLS  bool   `mapstructure:"tls" default:"false"`
	Worker WorkerConfig `mapstructure:"worker"`
}

type WorkerConfig struct {
	MaxConcurrentWorkflowTaskPollers int `mapstructure:"max_concurrent_workflow_task_pollers" default:"2"`
	MaxConcurrentWorkflowTaskExecutionSize int `mapstructure:"max_concurrent_workflow_task_execution_size" default:"10"`
	MaxConcurrentActivityTaskPollers int `mapstructure:"max_concurrent_activity_task_pollers" default:"2"`
	MaxConcurrentActivityExecutionSize int `mapstructure:"max_concurrent_activity_execution_size" default:"10"`
}

func (c *TemporalConfig) ToClientOptions() client.Options {
	opts := client.Options{
		HostPort: c.Addr,
	}

	if c.TLS {
		opts.ConnectionOptions = client.ConnectionOptions{
			TLS: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		}
	}

	return opts
}

type GitHubConfig struct {
	Token string `mapstructure:"token" default:"" validate:"required"`
}

type Config struct {
	GitHub GitHubConfig `mapstructure:"github"`
	LogLevel string        `mapstructure:"log_level" default:"info" validate:"required,oneof=debug info warn error"`
	HTTP     HTTPConfig    `mapstructure:"http"`
	Service  ServiceConfig `mapstructure:"service"`
	Temporal TemporalConfig `mapstructure:"temporal"`
}
