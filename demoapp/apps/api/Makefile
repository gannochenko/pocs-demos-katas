install:
	go mod download

run:
	@godotenv -f ./.env.local go run ./cmd/main.go

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

migrate:
	@godotenv -f ./.env.local go run ./script/migrate/migrate.go

seed:
	@godotenv -f ./.env.local go run ./script/seed/seed.go

generate_stubs:
	@openapi-generator generate -i ./openapi.yaml -g go-server -o ./go-server
