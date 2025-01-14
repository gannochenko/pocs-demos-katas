install:
	go mod download

run:
	@godotenv -f ./.env.local go run ./cmd/$(svc)/main.go

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

copy_test_env:
	@tr '\n' ';' < ./.env.test | pbcopy

new_migration:
	migrate create -ext sql -dir db/migrations $(name)

migrate_db:
	migrate -path db/migrations -database postgres://faceblur:faceblur@localhost:5432/faceblur?sslmode=disable up

seed_db:
	@godotenv -f ./.env.local go run ./script/seed/seed.go
