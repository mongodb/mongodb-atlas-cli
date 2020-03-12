# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

SOURCE_FILES?=./...
BINARY_NAME=mongocli

DESTINATION=./bin/${BINARY_NAME}
GOLANGCI_VERSION=v1.23.8
COVERAGE=coverage.out

VERSION=$(shell git describe --always --tags)
LINKER_FLAGS=-X github.com/mongodb/mongocli/internal/version.Version=${VERSION}

export PATH := ./bin:$(PATH)
export GO111MODULE := on

.PHONY: setup
setup:  ## Install dev tools
	@echo "==> Installing dependencies..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

.PHONY: link-git-hooks
link-git-hooks: ## Install git hooks
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: fmt
fmt: ## Format code
	@echo "==> Formatting all files..."
	find . -name '*.go' -not -path './third_party_notices*' -not -path './internal/mocks*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: test
test: ## Run tests
	@echo "==> Running tests..."
	go test -race -cover -count=1 -coverprofile ${COVERAGE} ./internal...

.PHONY: lint
lint: ## Run linter
	@echo "==> Linting all packages..."
	golangci-lint run $(SOURCE_FILES) -E goimports -E golint -E misspell -E unconvert -E maligned --skip-dirs ^internal/mocks/

.PHONY: check
check: test lint ## Run tests and linters

.PHONY: addlicense
addlicense:
	find . -name '*.go' -not -path './third_party_notices*' -not -path './internal/mocks*' | while read -r file; do addlicense -c "MongoDB Inc" "$$file"; done

.PHONY: gen-mocks
gen-mocks: ## Generate mocks
	@echo "==> Generating mocks"
	mockgen -source=internal/store/alert_configuration.go -destination=internal/mocks/mock_alert_configuration.go -package=mocks
	mockgen -source=internal/store/automation.go -destination=internal/mocks/mock_automation.go -package=mocks
	mockgen -source=internal/store/clusters.go -destination=internal/mocks/mock_clusters.go -package=mocks
	mockgen -source=internal/store/database_users.go -destination=internal/mocks/mock_database_users.go -package=mocks
	mockgen -source=internal/store/project_ip_whitelist.go -destination=internal/mocks/mock_project_ip_whitelist.go -package=mocks
	mockgen -source=internal/store/projects.go -destination=internal/mocks/mock_projects.go -package=mocks
	mockgen -source=internal/store/organizations.go -destination=internal/mocks/mock_organizations.go -package=mocks
	mockgen -source=internal/store/owners.go -destination=internal/mocks/mock_owners.go -package=mocks
	mockgen -source=internal/store/continuous_snapshots.go -destination=internal/mocks/mock_continuous_snapshots.go -package=mocks
	mockgen -source=internal/store/continuous_jobs.go -destination=internal/mocks/mock_continuous_jobs.go -package=mocks

.PHONY: build
build: ## Generate a binary in ./bin
	@echo "==> Building binary"
	go build -ldflags "${LINKER_FLAGS}" -o ${DESTINATION}

.PHONY: e2e-test
e2e-test: build ## Run E2E tests
	@echo "==> Running E2E tests..."
	# the target assumes the MCLI-* environment variables are exported
	go test -v -p 1 -parallel 1 -tags=e2e ./e2e...

.PHONY: install
install: ## Install a binary in $GOPATH/bin
	go install -ldflags "${LINKER_FLAGS}"

.PHONY: gen-notices
gen-notices: ## Generate 3rd party notices
	@echo "==> Generating 3rd party notices"
	@chmod -R 777 ./third_party_notices
	@rm -Rf third_party_notices
	go-licenses save "github.com/mongodb/mongocli" --save_path=third_party_notices

.PHONY: release
release: gen-notices ## Use goreleaser to generate builds and publish
	goreleaser --rm-dist

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
