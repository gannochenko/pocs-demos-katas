install:
	go mod download

run:
	@godotenv -f ./.env.local go run ./cmd/main.go

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

copy_test_env:
	@tr '\n' ';' < ./.env.test | pbcopy

seed:
	@godotenv -f ./.env.local go run ./script/seed/seed.go

generate_stubs:
	@openapi-generator generate -i ./openapi.yaml -g go-server -o ./go-server

new_migration:
	migrate create -ext sql -dir db/migrations $(name)

run_migrations:
	@godotenv -f ./.env.local $(MAKE) drop_db
	@godotenv -f ./.env.local $(MAKE) create_db
	@godotenv -f ./.env.local $(MAKE) migrate_db

run_test_infra_migration:
	@godotenv -f ./.env.local $(MAKE) drop_test_db
	@godotenv -f ./.env.local $(MAKE) create_test_db
	@godotenv -f ./.env.local $(MAKE) migrate_test_db

drop_db:
	psql "postgresql://demoapp:demoapp@localhost:5432/postgres" -c "DROP DATABASE IF EXISTS demoapp;"

drop_test_db:
	psql "postgresql://demoapp:demoapp@localhost:5432/postgres" -c "DROP DATABASE IF EXISTS demoapptest;"

create_db:
	psql "postgresql://demoapp:demoapp@localhost:5432/postgres" -c "CREATE DATABASE demoapp;"

create_test_db:
	psql "postgresql://demoapp:demoapp@localhost:5432/postgres" -c "CREATE DATABASE demoapptest;"

migrate_db:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapp?sslmode=disable up

migrate_test_db:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapptest?sslmode=disable up

migrate_db_up:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapp?sslmode=disable up 1

migrate_db_down:
	migrate -path db/migrations -database postgres://demoapp:demoapp@localhost:5432/demoapp?sslmode=disable down 1
