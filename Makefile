default: clean test build

GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard cmd/*.go)

export GO111MODULE = on

ARTIFACT_NAME = nhood-engine-utils

.PHONY: clean
clean:
	@echo "Cleaning:"
	go clean ./...
	@echo "...done"

.PHONY: install-dependencies
install-dependencies:
	@echo "Installing dependencies:"
	go mod vendor
	@echo "...done"

.PHONY: build
build: install-dependencies
	@echo "Building application:"
	@GOBIN=$(GOBIN) go install $(GOFILES)
	@echo "...done"

.PHONY: test
test: install-dependencies
	@echo "Running tests:"
	go test -v -cover ./pkg/...
	@echo "...done"

.PHONY: run
run: build
	./bin/engine-utils

.PHONY: release-ci
release-ci:
	@test $(GITHUB_USERNAME) || ( echo "GITHUB_USERNAME not set" & exit 1 )
	@test $(GITHUB_TOKEN) || ( echo "GITHUB_TOKEN not set" & exit 2 )
	@test $(GITHUB_EMAIL) || ( echo "GITHUB_EMAIL not set" & exit 3 )
	@test $(NEW_VERSION) || ( echo "NEW_VERSION not set" & exit 4 )
	@echo "Releasing maven artifacts [CI]:"
	git config --global user.email ${GITHUB_EMAIL} && \
	git config --global user.name ${GITHUB_USERNAME} && \
	git tag -a v${NEW_VERSION} -m "${NEW_VERSION}" && \
	git push --tags https://${GITHUB_TOKEN}@github.com/nhood-org/${ARTIFACT_NAME}.git master
	@echo "...done"

.PHONY: trigger-circle-ci-release
trigger-circle-ci-release:
	@test $(ARTIFACT_NAME) || ( echo "ARTIFACT_NAME not set" & exit 1 )
	@test $(CIRCLE_CI_USER_TOKEN) || ( echo "CIRCLE_CI_USER_TOKEN not set" & exit 2 )
	@test $(NEW_VERSION) || ( echo "NEW_VERSION not set" & exit 3 )
	@echo "Triggering docker release:"
	curl -u ${CIRCLE_CI_USER_TOKEN}: \
		-d build_parameters[CIRCLE_JOB]=release \
		-d build_parameters[VERSION]=${NEW_VERSION} \
		https://circleci.com/api/v1.1/project/github/nhood-org/${ARTIFACT_NAME}/tree/master
	@echo "...done"
