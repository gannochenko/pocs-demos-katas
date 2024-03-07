run:
	@godotenv -f ./.env.local go run ./cmd/main.go

lint:
	golangci-lint run --out-format=github-actions -v --timeout=10m0s --config ./.golangci-lint.yml

test:
	go test -short -mod=mod -v -p=1 -count=1 ./...

build_image:
	@./scripts/build-push-image.sh -a build -t $(tag) -e $(env)

push_image:
	@./scripts/build-push-image.sh -a push -t $(tag) -e $(env)
