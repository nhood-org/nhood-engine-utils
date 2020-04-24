GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard cmd/*.go)

export GO111MODULE = on

default: clean test build

clean:
	@echo "Cleaning:"
	go clean ./...
	@echo "...done"

install-dependencies:
	@echo "Installation of dependencies:"
	go mod vendor
	@echo "...done"

build: install-dependencies
	@echo "Installation:"
	@GOBIN=$(GOBIN) go install $(GOFILES)
	@echo "...done"

test: install-dependencies
	@echo "Running tests:"
	go test -v -cover ./pkg/...
	@echo "...done"

run: build
	./bin/engine-utils

.PHONY: clean test install run
