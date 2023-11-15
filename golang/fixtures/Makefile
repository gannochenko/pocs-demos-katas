run:
	@godotenv -f ./.env.local go run ./cmd/main.go

run_local_infra:
	@docker-compose -f ./infra/docker-compose-local.yml up --build

serialize_env_local:
	@tr '\n' ';' < ./.env.local

migrate:
	@godotenv -f ./.env.local go run ./script/migrate/migrate.go

seed:
	@godotenv -f ./.env.local go run ./script/seed/seed.go

cleanup:
	@godotenv -f ./.env.local go run ./script/cleanup/cleanup.go

create_database:
	@godotenv -f ./.env.local ./script/create_database.sh

copy_test_env:
	@tr '\n' ';' < ./.env.test | pbcopy