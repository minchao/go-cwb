SHELL := /bin/bash -o pipefail

.PHONY: help
## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test
## test: run unit tests
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
