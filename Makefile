.DEFAULT_GOAL := help
.PHONY: dist
GRC=$(shell which grc)

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

dist: ## Compile the app into a binary for Linux
dist:
	@CGO_ENABLED=0 GOOS=linux go build -o dist/bin/trailrcore main.go

deploy: ## Deploy the application through a docker image
deploy: dist
	./deploy.sh ghcr.io/arthureichelberger/trailrcore -f ./dist/Dockerfile ./dist
