run:
	@godotenv -f ./.env.local go run ./cmd/main.go

run_local_infra:
	@docker-compose -f ./infra/docker-compose-local.yml up --build
