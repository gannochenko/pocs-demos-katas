install:
	go mod download

run:
	@godotenv -f ./.env.local go run ./cmd/main.go

copy_env_local:
	@tr '\n' ';' < ./.env.local

migrate:
	@godotenv -f ./.env.local go run ./script/migrate/migrate.go

seed:
	@godotenv -f ./.env.local go run ./script/seed/seed.go
