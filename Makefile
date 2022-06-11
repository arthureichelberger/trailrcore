.DEFAULT_GOAL := help
.PHONY: dist
GRC=$(shell which grc)
IMG="ghcr.io/arthureichelberger/trailrcore"

-include make.properties

help: ## Makefile help
help:
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

start: ## Start the app locally
start:
	go run main.go

build: ## Compile the app into a binary for macOS
build:
	go build -o dist/bin/local/trailrcore main.go

unit: ## Run unit tests
unit:
	ENVIRONMENT=test $(GRC) go test -v -p=1 -count=1 -race -tags=unit ./... -timeout 2m

integration: ## Run integration tests
integration: tern migrate
	ENVIRONMENT=test $(GRC) go test -v -p=1 -count=1 -race -tags=integration ./... -timeout 2m

test: ## Run complete test suite
test:
	ENVIRONMENT=test $(GRC) go test -v -p=1 -count=1 -race -tags=unit,integration ./... -timeout 2m

dist: ## Compile the app into a binary for Linux
dist:
	@CGO_ENABLED=0 GOOS=linux go build -o dist/bin/trailrcore main.go

deploy: ## Deploy the application through a docker image
deploy: dist
	./deploy.sh $(IMG) -f ./dist/Dockerfile ./dist

docker: ## Compile the app and build it into a docker image locally
docker: dist
	docker build -t "${IMG}:local" -f ./dist/Dockerfile ./dist

coverage: ## Compute the code coverage
coverage: tern migrate
	./coverage.sh

tern: ## Install tern for database migrations
tern:
	go install github.com/jackc/tern@latest

migrate: ## Run migrations
migrate:
	tern migrate --config dist/migrations/tern.conf --migrations dist/migrations

dbreset: ## Reset the database
dbreset:
	docker compose rm -sfv db
	docker volume rm -f trailrcore_dbdata
	docker compose up -d
	sleep 2
