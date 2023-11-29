run:
	@godotenv -f ./.env.local go run ./cmd/main.go

run_local_infra:
	@docker-compose -f ./infra/docker-compose-local.yml up --build

serialize_env_local:
	@tr '\n' ';' < ./.env.local

build_lib:
	 @GOOS=js GOARCH=wasm go build -o build/main.wasm ./lib/main.go
