install:
	go mod download

run:
	@godotenv -f ./.env.local go run ./cmd/main.go

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

seed:
	@godotenv -f ./.env.local go run ./script/seed/seed.go

generate_stubs:
	@openapi-generator generate -i ./openapi.yaml -g go-server -o ./go-server

new_migration:
	migrate create -ext sql -dir db/migrations $(name)

run_migrations:
	@godotenv -f ./.env.local $(MAKE) drop_db
	@godotenv -f ./.env.local $(MAKE) create_db
	@godotenv -f ./.env.local $(MAKE) migrate_db serviceMigrationsPath=$(serviceMigrationsPath)

drop_db:
	psql -U demoapp -d postgres -h localhost -p 5432 -c 'DROP DATABASE IF EXISTS "demoapp"'

create_db:
	psql -U demoapp -d postgres -h localhost -p 5432 -c 'CREATE DATABASE "demoapp"'

migrate_db:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapp?sslmode=disable up

migrate_db_up:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapp?sslmode=disable up 1

migrate_db_down:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapp?sslmode=disable down 1
