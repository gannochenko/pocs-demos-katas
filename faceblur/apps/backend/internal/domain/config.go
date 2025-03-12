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

type RabbitMqConfigEventBus struct {
	QueueName    string `envconfig:"QUEUE_NAME"`
	RoutingKey   string `envconfig:"ROUTING_KEY"`
	ExchangeName string `envconfig:"EXCHANGE_NAME"`
}

type RabbitMqConfig struct {
	DSN      string                 `envconfig:"DSN"`
	EventBus RabbitMqConfigEventBus `envconfig:"EVENT_BUS"`
}

type BackendConfig struct {
	HTTP     HTTPConfig `envconfig:"HTTP"`
	GRPCPort int        `envconfig:"GRPC_PORT" default:"4646" desc:"grpc port"`
	Worker   WorkerConfig `envconfig:"WORKER"`
}

type WorkerConfig struct {
	ThreadCount int `envconfig:"THREAD_COUNT" default:"10" desc:"worker thread count"`
}

type Config struct {
	Postgres PostgresConfig `envconfig:"POSTGRES"`
	LogLevel string         `envconfig:"LOG_LEVEL" default:"info" desc:"logging level"`
	Auth0    Auth0Config    `envconfig:"AUTH0"`
	Env      string         `envconfig:"ENV"`
	GCP      GCPConfig      `envconfig:"GCP"`
	Storage  StorageConfig  `envconfig:"STORAGE"`
	RabbitMq RabbitMqConfig `envconfig:"RABBITMQ"`

	// app-specific
	Backend BackendConfig `envconfig:"BACKEND"`
}
