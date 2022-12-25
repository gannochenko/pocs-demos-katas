help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Common project related commands.
# ---------------------------------

install:
	@go mod download

create_resources:
	@psql -U test -d postgres -h localhost -p 54321 -c 'CREATE TABLE IF NOT EXISTS elements ( id char(36) primary key, title varchar(255) )'
	@psql -U test -d postgres -h localhost -p 54322 -c 'CREATE TABLE IF NOT EXISTS elements ( id char(36) primary key, title varchar(255) )'

run_infra:
	@docker-compose up

stop_infra:
	@docker-compose stop

run:
	@go run ./cmd/main.go
