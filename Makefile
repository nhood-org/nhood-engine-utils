#
# Before running any of the commands from within this Makefile,
# make sure your cloned repository is located in one of the sub-directories of the following path:
# $GOPATH/src/github.com/nhood-org/...
#

GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)

export GO111MODULE = on

default: clean test install

clean:
	@echo "Cleaning:"
	go clean
	@echo "...done"

install-dependencies:
	@echo "Installation of dependencies:"
	go mod vendor
	@echo "...done"

install: install-dependencies
	@echo "Installation:"
	@GOBIN=$(GOBIN) go install $(GOFILES)
	@echo "...done"

test: install-dependencies
	@echo "Running tests:"
	go test -v -cover ./pkg/...
	@echo "...done"

run:
	go run main.go

.PHONY: clean test install run
