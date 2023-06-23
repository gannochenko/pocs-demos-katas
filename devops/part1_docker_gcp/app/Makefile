run:
	@godotenv -f ./.env.local go run ./cmd/main.go

build_image:
	@docker build -t devops.app:part1 .

run_image:
	@docker run -p 8080:8080 devops.app:part1
