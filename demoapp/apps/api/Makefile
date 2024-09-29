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

generate_stubs:
	@openapi-generator generate -i ./openapi.yaml -g go-server -o ./go-server
