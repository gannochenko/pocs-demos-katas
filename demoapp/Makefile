help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	@$(MAKE) -C ./apps/api/ install
	@$(MAKE) -C ./apps/dashboard/ install

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

run: ## Run an application
ifeq ($(app),)
	$(error Please specify the "app" parameter. Example: "make run app=service")
else
	@$(MAKE) -C ./apps/$(app)/ run
endif
