include .env
export

.PHONY: test clean build up down run-ip run-token

clean: ## Clean all temp files
	@sudo rm -rf coverage*

test: ## Run unit-tests
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -html coverage.out -o coverage.html

build: ## Build the container image
	@docker build -t gustavo-villar/go-rate-limiter:dev -f Dockerfile .

up: ## Put the compose containers up
	@docker-compose up -d

down: ## Put the compose containers down
	@docker-compose down

run-ip:
	@bash run-ip.sh

run-token:
	@bash run-token.sh