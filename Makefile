APP_NAME   ?= gofiber-boilerplate
IMAGE_NAME ?= $(APP_NAME):latest
PORT       ?= 3000

.PHONY: help run build test fmt vet tidy docker-build docker-run

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

run: ## Run the app locally
	go run ./cmd

build: ## Build the binary into bin/
	go build -o bin/app ./cmd

test: ## Run the test suite
	go test ./...

fmt: ## Format all Go files
	gofmt -w .

vet: ## Run go vet
	go vet ./...

tidy: ## Tidy go.mod and go.sum
	go mod tidy

docker-build: ## Build the Docker image
	docker build -t $(IMAGE_NAME) .

docker-run: ## Run the image (requires .env: cp .env.example .env)
	docker run --rm -p $(PORT):3000 --env-file .env $(IMAGE_NAME)
