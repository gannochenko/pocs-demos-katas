run:
	@godotenv -f ./.env.local go run ./cmd/main.go

build_image:
	@docker build -t devops-app:$(tag) .

run_image:
	@docker run -p 8080:8080 devops-app:$(tag)

push_image:
    @docker push asia-east1-docker.pkg.dev/go-app-390716/devops/devops-app:$(tag)
