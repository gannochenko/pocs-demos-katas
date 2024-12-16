help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Variables.
# ---------------------------------

PROTO_DIR := ./protobuf
PROTO_OUT_DIR := .

# Find all .proto files recursively
PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	#@$(MAKE) -C ./apps/api/ install
	#@$(MAKE) -C ./apps/dashboard/ install
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	#@$(MAKE) -C ./apps/service_name/ create_resources

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

build_proto:
	@protoc \
      --go_out=$(OUT_DIR) \
	  --go-grpc_out=$(OUT_DIR) \
	  --go_opt=paths=source_relative \
	  --go-grpc_opt=paths=source_relative \
	  -I $(PROTO_DIR) \
      $(PROTO_FILES)

run: ## Run an application
ifeq ($(app),)
	$(error Please specify the "app" parameter. Example: "make run app=service")
else
	@$(MAKE) -C ./apps/$(app)/ run
endif
