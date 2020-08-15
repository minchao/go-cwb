SHELL := /bin/bash -o pipefail

.PHONY: help
## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: lint
## lint: run linters
lint:
	golangci-lint run

.PHONY: test
## test: run unit tests
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: cover
## cover: open a browser and displaying coverage profile
cover:
	go tool cover -html=coverage.txt
