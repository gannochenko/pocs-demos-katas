package domain

type PostgresConfig struct {
	DatabaseDSN string `envconfig:"DB_DSN" desc:"database connection DSN"`
}

type Config struct {
	Postgres PostgresConfig `envconfig:"POSTGRES"`
	LogLevel string         `envconfig:"LOG_LEVEL" default:"info" desc:"logging level"`
	HTTPPort int            `envconfig:"HTTP_PORT" default:"4545" desc:"http service level"`
}
