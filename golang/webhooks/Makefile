OPENAPI_DIR := ./contracts/openapi
OPENAPI_OUT_DIR := ./lib/openapi

SPECS := $(shell find $(OPENAPI_DIR) -name "*.yaml" -o -name "*.yml")
GENERATED := $(SPECS:.yaml=_gen.go)

bootstrap:
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install github.com/joho/godotenv/cmd/godotenv@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

run_gateway:
	cd apps/gateway && godotenv -f ../../env.local go run cmd/gateway/main.go

generate_gateway_openapi:
	@find $(OPENAPI_DIR) \( -name "*.yaml" -o -name "*.yml" \) -exec sh -c ' \
		for spec; do \
			dir=$$(dirname "$$spec"); \
			base=$$(basename "$$spec" | sed "s/\.[^.]*$$//"); \
			output="./apps/gateway/internal/http/v1/$${base}_gen.go"; \
			echo "Generating $$output from $$spec"; \
			oapi-codegen -package "$$base" -generate types,server,strict-server "$$spec" > "$$output"; \
		done \
	' sh {} +

run_infra:
	@docker compose up
