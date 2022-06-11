.DEFAULT_GOAL := help
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
	go build -o dist/bin/trailrcore main.go
