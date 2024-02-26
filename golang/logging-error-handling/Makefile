run:
	@godotenv -f ./.env.local go run ./cmd/main.go

run_local_infra:
	@docker-compose -f ./infra/docker-compose-local.yml up --build

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

migrate:
	@godotenv -f ./.env.local go run ./script/migrate/migrate.go

seed:
	@godotenv -f ./.env.local go run ./script/seed/seed.go

# CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
