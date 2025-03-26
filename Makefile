# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

GOLANGCI_VERSION=v2.0.1
COVERAGE?=coverage.out
export GOCOVERDIR?=$(abspath cov)

GIT_SHA?=$(shell git rev-parse HEAD)

ATLAS_SOURCE_FILES?=./cmd/atlas
ifeq ($(OS),Windows_NT)
    ATLAS_VERSION?=$(shell powershell -Command "(git describe --match 'atlascli/v*') -replace '.*v(.*)', '$$1'")
	ATLAS_BINARY_NAME=atlas.exe
else
    ATLAS_VERSION?=$(shell git describe --match "atlascli/v*" | cut -d "v" -f 2)
	ATLAS_BINARY_NAME=atlas
endif
ATLAS_DESTINATION=./bin/$(ATLAS_BINARY_NAME)
ATLAS_INSTALL_PATH="${GOPATH}/bin/$(ATLAS_BINARY_NAME)"

LOCALDEV_IMAGE?=docker.io/mongodb/mongodb-atlas-local
LINKER_FLAGS=-s -w -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version.GitCommit=${GIT_SHA} -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version.Version=${ATLAS_VERSION} -X github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options.LocalDevImage=${LOCALDEV_IMAGE}
ATLAS_E2E_BINARY?=../../bin/${ATLAS_BINARY_NAME}

DEBUG_FLAGS=all=-N -l

TEST_CMD?=go test
UNIT_TAGS?=unit,e2eSnap
E2E_TAGS?=e2e
E2E_TIMEOUT?=60m
E2E_PARALLEL?=1
E2E_EXTRA_ARGS?=
export UPDATE_SNAPSHOTS?=skip
export E2E_SKIP_CLEANUP?=false 

ifeq ($(OS),Windows_NT)
	export PATH := .\bin;$(shell go env GOPATH)\bin;$(PATH)
else
	export PATH := ./bin:$(shell go env GOPATH)/bin:$(PATH)
endif
export TERM := linux-m
export GO111MODULE := on
export GOTOOLCHAIN := local
export ATLAS_E2E_BINARY

.PHONY: pre-commit
pre-commit:  ## Run pre-commit hook
	@echo "==> Running pre-commit hook..."
	@scripts/pre-commit.sh

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
setup: deps devtools link-git-hooks ## Set up dev env

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
test: unit-test

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: fix-lint
fix-lint: ## Fix linting errors
	golangci-lint run --fix

.PHONY: check
check: test fix-lint ## Run tests and linters

.PHONY: check-templates
check-templates: ## Verify templates
	go run ./tools/cmd/templates-checker

.PHONY: addcopy
addcopy: ## Add missing license to files
	@scripts/add-copy.sh

.PHONY: generate
generate: gen-docs gen-mocks gen-api-commands ## Generate docs, mocks, code, api commands, all auto generated assets

.PHONY: apply-overlay
apply-overlay: ## Apply overlay on openapi spec
	@echo "==> Applying overlay"
	go run ./tools/cmd/apply-overlay --spec ./tools/internal/specs/spec.yaml --overlay ./tools/internal/specs/overlays/\*.yaml > ./tools/internal/specs/spec-with-overlays.yaml

.PHONY: gen-api-commands
gen-api-commands: apply-overlay ## Generate api commands
	@echo "==> Generating api commands"
	go run ./tools/cmd/api-generator --spec ./tools/internal/specs/spec-with-overlays.yaml --output-type commands > ./internal/api/commands.go

.PHONY: gen-docs-metadata
gen-docs-metadata: apply-overlay ## Generate docs metadata
	@echo "==> Generating docs metadata"
	go run ./tools/cmd/api-generator --spec ./tools/internal/specs/spec-with-overlays.yaml --output-type metadata > ./tools/cmd/docs/metadata.go

.PHONY: otel
otel: ## Generate code
	go run ./tools/cmd/otel $(SPAN) --attr $(ATTRS)

.PHONY: gen-mocks
gen-mocks: ## Generate mocks
	@echo "==> Generating mocks"
	rm -rf ./internal/mocks
	go generate ./internal...

.PHONY: gen-docs
gen-docs: gen-docs-metadata ## Generate docs for atlascli commands
	@echo "==> Generating docs"
	go run -ldflags "$(LINKER_FLAGS)" ./tools/cmd/docs

.PHONY: build
build: ## Generate an atlas binary in ./bin
	@echo "==> Building $(ATLAS_BINARY_NAME) binary"
	go build -ldflags "$(LINKER_FLAGS)" $(BUILD_FLAGS) -o $(ATLAS_DESTINATION) $(ATLAS_SOURCE_FILES)

.PHONY: build-debug
build-debug: ## Generate a binary in ./bin for debugging atlascli
	@echo "==> Building $(ATLAS_BINARY_NAME) binary for debugging"
	go build -gcflags="$(DEBUG_FLAGS)" -ldflags "$(LINKER_FLAGS)" $(BUILD_FLAGS) -cover -o $(ATLAS_DESTINATION) $(ATLAS_SOURCE_FILES)

.PHONY: e2e-test
e2e-test: build-debug ## Run E2E tests
# the target assumes the MCLI_* environment variables are exported
	@echo "==> Running E2E tests..."
	$(TEST_CMD) -v -p 1 -parallel $(E2E_PARALLEL) -v -timeout $(E2E_TIMEOUT) -tags="$(E2E_TAGS)" ./test/e2e... $(E2E_EXTRA_ARGS)

.PHONY: unit-test
unit-test: ## Run unit-tests
unit-test: export DO_NOT_TRACK=1
unit-test: export E2E_SKIP_CLEANUP=true
unit-test: export UPDATE_SNAPSHOTS=false
unit-test: export MONGODB_ATLAS_ORG_ID=5f0f5b3e0f2912c8b8f3b9b9
unit-test: export MONGODB_ATLAS_PROJECT_ID=5f0f5b3e0f2912c8b8f3b9b9
unit-test: export MONGODB_ATLAS_PRIVATE_API_KEY=12345678-abcd-ef01-2345-6789abcdef01
unit-test: export MONGODB_ATLAS_PUBLIC_API_KEY=ABCDEF01
unit-test: export MONGODB_ATLAS_OPS_MANAGER_URL=http://localhost:8080
unit-test: export MONGODB_ATLAS_SERVICE=cloud
unit-test: export IDENTITY_PROVIDER_ID=5f0f5b3e0f2912c8b8f3b9b9
unit-test: export E2E_CLOUD_ROLE_ID=5f0f5b3e0f2912c8b8f3b9b9
unit-test: export E2E_TEST_BUCKET=test-bucket
unit-test: export E2E_FLEX_INSTANCE_NAME=instance_name
unit-test: build-debug
	@echo "==> Running unit tests..."
	$(TEST_CMD) -parallel $(E2E_PARALLEL) --tags="$(UNIT_TAGS)" -cover -coverprofile $(COVERAGE) -count=1 ./...

.PHONY: install
install: ## Install a binary in $GOPATH/bin
	@echo "==> Installing $(ATLAS_BINARY_NAME) to $(ATLAS_INSTALL_PATH)"
	go install -ldflags "$(LINKER_FLAGS)" $(ATLAS_SOURCE_FILES)
	@echo "==> Done..."

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: check-library-owners
check-library-owners: ## Check that all the dependencies in go.mod has a owner in library_owners.json
	@echo "==> Check library_owners.json"
	go run ./tools/cmd/libraryowners/main.go
	./scripts/verify-library-owners-sorted.sh

.PHONY: update-atlas-sdk
update-atlas-sdk: ## Update the atlas-sdk dependency
	./scripts/update-sdk.sh

.PHONY: update-openapi-spec
update-openapi-spec: ## Update the openapi spec
	./scripts/update-openapi-spec.sh

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
