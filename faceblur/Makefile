help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Variables.
# ---------------------------------

PROTO_DIR := ./protobuf
PROTO_OUT_BACKEND_DIR := ./apps/backend/proto
PROTO_OUT_DASHBOARD_DIR := ./apps/dashboard/src/proto

# Find all .proto files recursively
PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto" ! -path "$(PROTO_DIR)/google/*")

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	#@$(MAKE) -C ./apps/api/ install
	#@$(MAKE) -C ./apps/dashboard/ install
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@npm install @gannochenko/protoweb -g

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	@godotenv -f ./.env.local python3.7 ./scripts/bootstrap_gcp.py

list_resources: ## List all local resources
	#@$(MAKE) -C ./apps/service_name/ list_resources

remove_resources: ## Remove all local resources

run_local_infra: ## Run the infrastructure locally
	@docker compose up

run_test_infra: ## Run the infrastructure for testing
	docker compose -f ./docker compose-test.yml down
	docker compose -f ./docker compose-test.yml up

run_test_infra_migration:
	@$(MAKE) -C ./apps/api/ run_test_infra_migration

#stop_infra: ## Run the infrastructure locally
#	@docker compose stop

build_backend_proto:
	@rm -rf $(PROTO_OUT_BACKEND_DIR)
	@mkdir -p $(PROTO_OUT_BACKEND_DIR)
	@protoc \
      --go_out=$(PROTO_OUT_BACKEND_DIR) \
	  --go-grpc_out=$(PROTO_OUT_BACKEND_DIR) \
	  --go_opt=paths=source_relative \
	  --go-grpc_opt=paths=source_relative \
	  --grpc-gateway_out=paths=source_relative:$(PROTO_OUT_BACKEND_DIR) \
	  -I $(PROTO_DIR) \
      $(PROTO_FILES)

build_dashboard_proto:
	@rm -rf $(PROTO_OUT_DASHBOARD_DIR)
	@mkdir -p $(PROTO_OUT_DASHBOARD_DIR)
	@protoweb build -i $(PROTO_DIR) -o $(PROTO_OUT_DASHBOARD_DIR) -t ./apps/dashboard/protoweb.cjs

run: ## Run an application
ifeq ($(app),)
	$(error Please specify the "app" parameter. Example: "make run app=service svc=api")
else
	@$(MAKE) -C ./apps/$(app)/ run svc=$(svc)
endif
