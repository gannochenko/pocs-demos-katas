help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	@pip install google-api-python-client
	@pip install google.cloud.storage
	@pip install google-cloud
	@pip install google-cloud-vision

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	@godotenv -f ./.env.local python ./scripts/bootstrap-storage.py

run_infra: ## Run the infrastructure locally
	@docker-compose up

stop_infra: ## Run the infrastructure locally
	@docker-compose stop
