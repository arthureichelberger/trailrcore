#!/bin/sh

set -e

# Run tests and extract code coverage
ENVIRONMENT=test go test -v -p=1 -count=1 -race -tags=unit,integration ./... -timeout 2m -coverprofile coverage.out -covermode atomic
go tool cover -func coverage.out > /dev/null
TOTAL_COVERAGE=$(go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+')

echo "TOTAL_COVERAGE=$TOTAL_COVERAGE" >> $GITHUB_ENV
