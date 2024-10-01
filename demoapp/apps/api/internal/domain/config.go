package domain

type PostgresConfig struct {
	DatabaseDSN string `envconfig:"DB_DSN" desc:"database connection DSN"`
}

type Auth0Config struct {
	Audience string `envconfig:"AUDIENCE"`
	Domain   string `envconfig:"DOMAIN"`
}

type Config struct {
	Postgres PostgresConfig `envconfig:"POSTGRES"`
	LogLevel string         `envconfig:"LOG_LEVEL" default:"info" desc:"logging level"`
	HTTPPort int            `envconfig:"HTTP_PORT" default:"4545" desc:"http service level"`
	Auth0    Auth0Config    `envconfig:"AUTH0"`
}
