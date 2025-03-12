help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Variables.
# ---------------------------------

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	@$(MAKE) -C ./apps/monitor/ install

install_scripts:
	@cd ./scripts; python3 -m venv venv; source venv/bin/activate; pip install -r requirements.txt; deactivate

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	@godotenv -f ./.env.local python3 ./scripts/bootstrap_gcp.py
	@godotenv -f ./.env.local python3 ./scripts/bootstrap_eventbus.py

create_app_resources:
	@$(MAKE) -C ./apps/backend/ create_resources

run_local_infra: ## Run the infrastructure locally
	@docker compose up

copy_local_env:
	@tr '\n' ';' < ./.env.local | pbcopy

run: ## Run an application
ifeq ($(app),)
	$(error Please specify the "app" parameter. Example: "make run app=service svc=api")
else
	@$(MAKE) -C ./apps/$(app)/ run svc=$(svc)
endif
