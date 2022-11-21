help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	#@$(MAKE) -C ./apps/service_name/ install

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	#@$(MAKE) -C ./apps/service_name/ create_resources

list_resources: ## List all local resources
	#@$(MAKE) -C ./apps/service_name/ list_resources

remove_resources: ## Remove all local resources

seed_database:
	#@$(MAKE) -C ./apps/service_name/ seed_database

migrate_databases: ## Migrate databases of all applications
	#@$(MAKE) -C ./apps/service_name/ migrate_database

run_infra: ## Run the infrastructure locally
	@docker-compose up

stop_infra: ## Run the infrastructure locally
	@docker-compose stop

run: ## Run an application
ifeq ($(app),)
	$(error Please specify the "app" parameter. Example: "make run app=service")
else
	@$(MAKE) -C ./apps/$(app)/ run
endif

run_docs: ## Run the documentation locally
	@mkdocs serve -f docs/mkdocs.yml
