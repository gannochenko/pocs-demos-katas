package domain

type PostgresConfig struct {
	DatabaseDSN string `envconfig:"DB_DSN" desc:"database connection DSN"`
}

type Auth0Config struct {
	Audience string `envconfig:"AUDIENCE"`
	Domain   string `envconfig:"DOMAIN"`
}

type GCPConfig struct {
	ServiceAccount string `envconfig:"SERVICE_ACCOUNT"`
}

type StorageConfig struct {
	ImageBucketName string `envconfig:"IMAGE_BUCKET_NAME"`
}

type CorsConfig struct {
	Origin []string `envconfig:"ORIGIN" default:"" desc:"http cors origin"`
}

type HTTPConfig struct {
	Port int        `envconfig:"PORT" default:"4545" desc:"http port"`
	Cors CorsConfig `envconfig:"CORS"`
}

type RabbitMqConfig struct {
	Host      string `envconfig:"HOST"`
	Port      int    `envconfig:"PORT"`
	QueueName string `envconfig:"QUEUE_NAME"`
}

type Config struct {
	Postgres PostgresConfig `envconfig:"POSTGRES"`
	LogLevel string         `envconfig:"LOG_LEVEL" default:"info" desc:"logging level"`
	HTTP     HTTPConfig     `envconfig:"HTTP"`
	GRPCPort int            `envconfig:"GRPC_PORT" default:"4646" desc:"grpc port"`
	Auth0    Auth0Config    `envconfig:"AUTH0"`
	Env      string         `envconfig:"ENV"`
	GCP      GCPConfig      `envconfig:"GCP"`
	Storage  StorageConfig  `envconfig:"STORAGE"`
	RabbitMq RabbitMqConfig `envconfig:"RABBITMQ"`
}
