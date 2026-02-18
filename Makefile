# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

GOLANGCI_VERSION=v2.10.1
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

DEBUG_FLAGS=all=-N -l

TEST_CMD?=go test
E2E_TEST_PACKAGES?=./test/e2e/...
E2E_TIMEOUT?=60m
E2E_PARALLEL?=1
E2E_EXTRA_ARGS?=
export TEST_MODE?=live

ifeq ($(OS),Windows_NT)
	export PATH := .\bin;$(shell go env GOPATH)\bin;$(PATH)
else
	export PATH := ./bin:$(shell go env GOPATH)/bin:$(PATH)
endif
export TERM := linux-m
export GO111MODULE := on
export GOTOOLCHAIN := local

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
	golangci-lint fmt

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
generate: gen-docs gen-mocks gen-api-commands gen-purls ## Generate docs, mocks, code, api commands, all auto generated assets

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

.PHONY: gen-mocks
gen-mocks: ## Generate mocks
	@echo "==> Generating mocks"
	rm -rf ./internal/mocks
	go generate ./internal...

.PHONY: gen-docs
gen-docs: gen-docs-metadata ## Generate docs for atlascli commands
	@echo "==> Generating docs"
	go run -ldflags "$(LINKER_FLAGS)" ./tools/cmd/docs

.PHONY: gen-purls
gen-purls: # Generate purls
	@echo "==> Generating Linux purls"
	GOOS=linux GOARCH=amd64 go build -trimpath -mod=readonly -o bin/atlas-linux ./cmd/atlas
	go version -m ./bin/atlas-linux | \
		awk '$$1 == "dep" || $$1 == "=>" { print "pkg:golang/" $$2 "@" $$3 }' | \
		LC_ALL=C sort > build/package/purls-linux.txt

	@echo "==> Generating Darwin purls"
	GOOS=darwin GOARCH=arm64 go build -trimpath -mod=readonly -o bin/atlas-darwin ./cmd/atlas
	go version -m ./bin/atlas-darwin | \
		awk '$$1 == "dep" || $$1 == "=>" { print "pkg:golang/" $$2 "@" $$3 }' | \
		LC_ALL=C sort > build/package/purls-darwin.txt

	@echo "==> Generating Windows purls"
	GOOS=windows GOARCH=amd64 go build -trimpath -mod=readonly -o bin/atlas-win ./cmd/atlas
	go version -m ./bin/atlas-win | \
		awk '$$1 == "dep" || $$1 == "=>" { print "pkg:golang/" $$2 "@" $$3 }' | \
		LC_ALL=C sort > build/package/purls-win.txt

	@echo "==> Merging purls"
	cat build/package/purls-linux.txt build/package/purls-darwin.txt build/package/purls-win.txt |  LC_ALL=C sort | uniq > build/package/purls.txt
	rm -rf build/package/purls-linux.txt build/package/purls-darwin.txt build/package/purls-win.txt

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
	$(TEST_CMD) -v -p 1 -parallel $(E2E_PARALLEL) -v -timeout $(E2E_TIMEOUT) ${E2E_TEST_PACKAGES} $(E2E_EXTRA_ARGS)
	go tool covdata textfmt -i $(GOCOVERDIR) -o $(COVERAGE)

.PHONY: e2e-test-snapshots
e2e-test-snapshots: build-debug ## Run E2E tests
	TEST_MODE=replay DO_NOT_TRACK=1 $(TEST_CMD) -v -timeout $(E2E_TIMEOUT) ${E2E_TEST_PACKAGES} $(E2E_EXTRA_ARGS)
	go tool covdata textfmt -i $(GOCOVERDIR) -o $(COVERAGE)

.PHONY: unit-test
unit-test: build-debug ## Run unit-tests
	@echo "==> Running unit tests..."
	$(TEST_CMD) -parallel $(E2E_PARALLEL) -short -cover -coverprofile $(COVERAGE) -count=1 ./...

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

.PHONY: update-atlas-cli-core
update-atlas-cli-core: ## Update the atlas-cli-core dependency to the latest SHA
	@echo "==> Updating atlas-cli-core to latest SHA..."
	go get github.com/mongodb/atlas-cli-core@master
	go mod tidy

.PHONY: update-openapi-spec
update-openapi-spec: ## Update the openapi spec
	./scripts/update-openapi-spec.sh

.PHONY: add-e2e-profiles
add-e2e-profiles: build ## Add e2e profiles
	./scripts/add-e2e-profiles.sh

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
