OPENAPI_DIR := ./contracts/openapi
OPENAPI_OUT_DIR := ./apps/api/internal/http

SPECS := $(shell find $(OPENAPI_DIR) -name "*.yaml" -o -name "*.yml")
GENERATED := $(SPECS:.yaml=_gen.go)

bootstrap:
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install github.com/joho/godotenv/cmd/godotenv@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

run_worker:
	@cd apps/worker && godotenv -f ../../env.local go run cmd/worker/main.go

migrate_worker:
	@cd apps/worker && godotenv -f ../../env.local go run ./cmd/migrate/main.go

run_api:
	@cd apps/api && godotenv -f ../../env.local go run cmd/api/main.go

generate_api_openapi:
	@rm -rf $(OPENAPI_OUT_DIR)/v1
	@mkdir -p $(OPENAPI_OUT_DIR)/v1
	@find $(OPENAPI_DIR) \( -name "*.yaml" -o -name "*.yml" \) -exec sh -c ' \
		for spec; do \
			dir=$$(dirname "$$spec"); \
			base=$$(basename "$$spec" | sed "s/\.[^.]*$$//"); \
			output="$(OPENAPI_OUT_DIR)/v1/$${base}_gen.go"; \
			echo "Generating $$output from $$spec"; \
			oapi-codegen -package "$$base" -generate types,server,strict-server "$$spec" > "$$output"; \
		done \
	' sh {} +

run_infra:
	@docker compose up
