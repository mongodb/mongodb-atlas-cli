# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

GOLANGCI_VERSION=v1.54.2
COVERAGE=coverage.out

MCLI_SOURCE_FILES?=./cmd/mongocli
MCLI_BINARY_NAME=mongocli
MCLI_VERSION?=$(shell git describe --match "mongocli/v*" | cut -d "v" -f 2)
MCLI_GIT_SHA?=$(shell git rev-parse HEAD)
MCLI_DESTINATION=./bin/$(MCLI_BINARY_NAME)
MCLI_INSTALL_PATH="${GOPATH}/bin/$(MCLI_BINARY_NAME)"
MCLI_E2E_BINARY?=../../../bin/${MCLI_BINARY_NAME}

ATLAS_SOURCE_FILES?=./cmd/atlas
ATLAS_BINARY_NAME=atlas
ATLAS_VERSION?=$(shell git describe --match "atlascli/v*" | cut -d "v" -f 2)
ATLAS_DESTINATION=./bin/$(ATLAS_BINARY_NAME)
ATLAS_INSTALL_PATH="${GOPATH}/bin/$(ATLAS_BINARY_NAME)"

LINKER_FLAGS=-s -w -X github.com/mongodb/mongodb-atlas-cli/internal/version.GitCommit=${MCLI_GIT_SHA}
MCLI_LINKER_FLAGS=${LINKER_FLAGS} -X github.com/mongodb/mongodb-atlas-cli/internal/config.ToolName=$(MCLI_BINARY_NAME) -X github.com/mongodb/mongodb-atlas-cli/internal/version.Version=${MCLI_VERSION}
ATLAS_LINKER_FLAGS=${LINKER_FLAGS} -X github.com/mongodb/mongodb-atlas-cli/internal/config.ToolName=atlascli -X github.com/mongodb/mongodb-atlas-cli/internal/version.Version=${ATLAS_VERSION}
ATLAS_E2E_BINARY?=../../../bin/${ATLAS_BINARY_NAME}

DEBUG_FLAGS=all=-N -l

TEST_CMD?=go test
UNIT_TAGS?=unit
INTEGRATION_TAGS?=integration
E2E_TAGS?=e2e
E2E_TIMEOUT?=60m
E2E_PARALLEL?=1

export PATH := $(shell go env GOPATH)/bin:$(PATH)
export PATH := ./bin:$(PATH)
export TERM := linux-m
export GO111MODULE := on
export MCLI_E2E_BINARY
export ATLAS_E2E_BINARY

.PHONY: deps
deps:  ## Download go module dependencies
	@echo "==> Installing go.mod dependencies..."
	go mod download
	go mod tidy

.PHONY: devtools
devtools:  ## Install dev tools
	@echo "==> Installing dev tools..."
	go install github.com/google/addlicense@latest
	go install github.com/golang/mock/mockgen@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/google/go-licenses@latest
	go install mvdan.cc/sh/v3/cmd/shfmt@latest
	go install github.com/icholy/gomajor@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)

.PHONY: setup
setup: deps devtools ## Set up dev env

.PHONY: link-git-hooks
link-git-hooks: ## Install git hooks
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: fmt
fmt: ## Format changed go
	@scripts/fmt.sh

.PHONY: fmt-all
fmt-all: ### Format all go files with goimports and gofmt
	find . -name "*.go" -not -path "./vendor/*" -not -path "./internal/mocks" -exec gofmt -w "{}" \;
	find . -name "*.go" -not -path "./vendor/*" -not -path "./internal/mocks" -exec goimports -l -w "{}" \;

.PHONY: test
test: unit-test integration-test fuzz-normalizer-test

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: fix-lint
fix-lint: ## Fix linting errors
	golangci-lint run --fix

.PHONY: check
check: test fix-lint ## Run tests and linters

.PHONY: addcopy
addcopy:
	@scripts/add-copy.sh

.PHONY: generate
generate: gen-docs gen-mocks gen-code ## Generate docs, mocks, code, all auto generated assets

.PHONY: gen-code
gen-code: ## Generate code
	@echo "==> Generating code"
	go run ./tools/cli-generator

.PHONY: gen-mocks
gen-mocks: ## Generate mocks
	@echo "==> Generating mocks"
	go generate ./internal...

.PHONY: gen-docs
gen-docs: gen-docs-mongocli gen-docs-atlascli ## Generate docs for commands

.PHONY: gen-docs-mongocli
gen-docs-mongocli: ## Generate docs for mongocli commands
	@echo "==> Generating docs for mongocli"
	go run -ldflags "$(MCLI_LINKER_FLAGS)" ./tools/mongoclidocs/main.go

.PHONY: gen-docs-atlascli
gen-docs-atlascli: ## Generate docs for atlascli commands
	@echo "==> Generating docs for atlascli"
	go run -ldflags "$(ATLAS_LINKER_FLAGS)" ./tools/atlasclidocs/main.go

.PHONY: build
build: build-mongocli ## Generate a binary for mongocli

.PHONY: build-all
build-all: build-mongocli build-atlascli ## Generate a binary for both CLIs

.PHONY: build-mongocli
build-mongocli: ## Generate a mongocli binary in ./bin
	@echo "==> Building $(MCLI_BINARY_NAME) binary"
	go build -ldflags "$(MCLI_LINKER_FLAGS)" -o $(MCLI_DESTINATION) $(MCLI_SOURCE_FILES)

.PHONY: build-atlascli
build-atlascli: ## Generate a atlascli binary in ./bin
	@echo "==> Building $(ATLAS_BINARY_NAME) binary"
	go build -ldflags "$(ATLAS_LINKER_FLAGS)" -o $(ATLAS_DESTINATION) $(ATLAS_SOURCE_FILES)

.PHONY: build-debug
build-debug: build-mongocli-debug build-atlascli-debug ## Generate binaries in ./bin for debugging both CLIs

.PHONY: build-mongocli-debug
build-mongocli-debug: ## Generate a binary in ./bin for debugging mongocli
	@echo "==> Building $(MCLI_BINARY_NAME) binary for debugging"
	go build -gcflags="$(DEBUG_FLAGS)" -ldflags "$(MCLI_LINKER_FLAGS)" -o $(MCLI_DESTINATION) $(MCLI_SOURCE_FILES)

.PHONY: build-atlascli-debug
build-atlascli-debug: ## Generate a binary in ./bin for debugging atlascli
	@echo "==> Building $(ATLAS_BINARY_NAME) binary for debugging"
	go build -gcflags="$(DEBUG_FLAGS)" -ldflags "$(ATLAS_LINKER_FLAGS)" -o $(ATLAS_DESTINATION) $(ATLAS_SOURCE_FILES)

.PHONY: e2e-test
e2e-test: build-all ## Run E2E tests
	@echo "==> Running E2E tests..."
	# the target assumes the MCLI_* environment variables are exported
	$(TEST_CMD) -v -p 1 -parallel $(E2E_PARALLEL) -timeout $(E2E_TIMEOUT) -tags="$(E2E_TAGS)" ./test/e2e...

.PHONY: integration-test
integration-test: ## Run integration tests
	@echo "==> Running integration tests..."
	$(TEST_CMD) --tags="$(INTEGRATION_TAGS)" -count=1 ./internal...

.PHONY: fuzz-normalizer-test
fuzz-normalizer-test: ## Run fuzz test
	@echo "==> Running fuzz test..."
	$(TEST_CMD) -fuzz=Fuzz -fuzztime 50s --tags="$(UNIT_TAGS)" -race ./internal/kubernetes/operator/resources

.PHONY: unit-test
unit-test: ## Run unit-tests
	@echo "==> Running unit tests..."
	$(TEST_CMD) --tags="$(UNIT_TAGS)" -race -cover -count=1 -coverprofile $(COVERAGE) ./...

.PHONY: install
install: install-mongocli install-atlascli ## Install binaries in $GOPATH/bin for both CLIs

.PHONY: install-mongocli
install-mongocli: ## Install mongocli binary in $GOPATH/bin
	@echo "==> Installing $(MCLI_BINARY_NAME) to $(MCLI_INSTALL_PATH)"
	go install -ldflags "$(MCLI_LINKER_FLAGS)" $(MCLI_SOURCE_FILES)
	@echo "==> Done..."

.PHONY: install-atlascli
install-atlascli: ## Install atlascli binary in $GOPATH/bin
	@echo "==> Installing $(ATLAS_BINARY_NAME) to $(ATLAS_INSTALL_PATH)"
	go install -ldflags "$(ATLAS_LINKER_FLAGS)" $(ATLAS_SOURCE_FILES)
	@echo "==> Done..."

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: check-library-owners
check-library-owners: ## Check that all the dependencies in go.mod has a owner in library_owners.json
	@echo "==> Check library_owners.json"
	go run ./tools/libraryowners/main.go

.PHONY: update-atlas-sdk
update-atlas-sdk: ## Update the atlas-sdk dependency
	$(eval LATEST_SDK_RELEASE := $(shell curl -sSfL -X GET  https://api.github.com/repos/mongodb/atlas-sdk-go/releases/latest | jq -r '.tag_name' | cut -d '.' -f 1 ))
	@echo "==> Updating SDK to latest major version $(LATEST_SDK_RELEASE)"
	gomajor get go.mongodb.org/atlas-sdk/$(LATEST_SDK_RELEASE)@latest
	go mod tidy
	 @echo "==> Done, remember to update build/ci/library_owners.json"

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
