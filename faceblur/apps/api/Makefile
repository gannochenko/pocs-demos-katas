install:
	go mod download

run:
	@godotenv -f ./.env.local go run ./cmd/main.go

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

copy_test_env:
	@tr '\n' ';' < ./.env.test | pbcopy
