package domain

type HTTPConfig struct {
	Addr string `envconfig:"ADDR" default:":8080"`
}

type ServiceConfig struct {
	Name    string `envconfig:"NAME" default:"tasks-service"`
	Version string `envconfig:"VERSION" default:"1.0.0"`
}

type Config struct {
	LogLevel string         `envconfig:"LOG_LEVEL" default:"info" desc:"logging level"`
	HTTP     HTTPConfig     `envconfig:"HTTP"`
	Service  ServiceConfig  `envconfig:"SERVICE"`
}
