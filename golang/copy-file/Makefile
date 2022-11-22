help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Common project related commands.
# ---------------------------------

install:
	@pip install google-api-python-client
	@pip install google.cloud.storage
	@pip install google-cloud
	@pip install google-cloud-vision
	@go mod download

create_resources:
	@godotenv -f ./.env.local python ./scripts/bootstrap-storage.py

run_infra:
	@docker-compose up

stop_infra:
	@docker-compose stop

run:
	@godotenv -f ./.env.local go run ./cmd/main.go
