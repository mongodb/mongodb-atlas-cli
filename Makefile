# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

SOURCE_FILES?=./...
BINARY_NAME=mongocli

DESTINATION=./bin/${BINARY_NAME}
GOLANGCI_VERSION=v1.27.0
COVERAGE=coverage.out
VERSION=$(shell git describe --always --tags)
LINKER_FLAGS=-X github.com/mongodb/mongocli/internal/version.Version=${VERSION}

E2E_CMD?=go test
E2E_TAGS?=e2e

export PATH := ./bin:$(PATH)
export GO111MODULE := on

.PHONY: setup
setup:  ## Install dev tools
	@echo "==> Installing dependencies..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)

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
	golangci-lint run $(SOURCE_FILES)

.PHONY: fix-lint
fix-lint: ## Fix lint errors
	@echo "==> Fixing lint errors"
	golangci-lint run $(SOURCE_FILES) --fix

.PHONY: check
check: test fix-lint ## Run tests and linters

.PHONY: addlicense
addlicense:
	find . -name '*.go' -not -path './third_party_notices*' -not -path './internal/mocks*' | while read -r file; do addlicense -c "MongoDB Inc" "$$file"; done

.PHONY: gen-mocks
gen-mocks: ## Generate mocks
	@echo "==> Generating mocks"
	mockgen -source=internal/config/profile.go -destination=internal/mocks/mock_profile.go -package=mocks
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
	mockgen -source=internal/store/agents.go -destination=internal/mocks/mock_agents.go -package=mocks
	mockgen -source=internal/store/checkpoints.go -destination=internal/mocks/mock_checkpoints.go -package=mocks
	mockgen -source=internal/store/alerts.go -destination=internal/mocks/mock_alerts.go -package=mocks
	mockgen -source=internal/store/global_alerts.go -destination=internal/mocks/mock_global_alerts.go -package=mocks
	mockgen -source=internal/store/events.go -destination=internal/mocks/mock_events.go -package=mocks
	mockgen -source=internal/store/process_measurements.go -destination=internal/mocks/mock_process_measurements.go -package=mocks
	mockgen -source=internal/store/process_disks.go -destination=internal/mocks/mock_process_disks.go -package=mocks
	mockgen -source=internal/store/process_disk_measurements.go -destination=internal/mocks/mock_process_disk_measurements.go -package=mocks
	mockgen -source=internal/store/process_databases.go -destination=internal/mocks/mock_process_databases.go -package=mocks
	mockgen -source=internal/store/host_measurements.go -destination=internal/mocks/mock_host_measurements.go -package=mocks
	mockgen -source=internal/store/indexes.go -destination=internal/mocks/mock_indexes.go -package=mocks
	mockgen -source=internal/store/processes.go -destination=internal/mocks/mock_processes.go -package=mocks
	mockgen -source=internal/store/logs.go -destination=internal/mocks/mock_logs.go -package=mocks
	mockgen -source=internal/store/hosts.go -destination=internal/mocks/mock_hosts.go -package=mocks
	mockgen -source=internal/store/host_databases.go -destination=internal/mocks/mock_host_databases.go -package=mocks
	mockgen -source=internal/store/host_disks.go -destination=internal/mocks/mock_host_disks.go -package=mocks
	mockgen -source=internal/store/host_disk_measurements.go -destination=internal/mocks/mock_host_disk_measurements.go -package=mocks
	mockgen -source=internal/store/diagnose_archive.go -destination=internal/mocks/mock_diagnose_archive.go -package=mocks

.PHONY: build
build: ## Generate a binary in ./bin
	@echo "==> Building binary"
	go build -ldflags "${LINKER_FLAGS}" -o ${DESTINATION}

.PHONY: e2e-test
e2e-test: build ## Run E2E tests
	@echo "==> Running E2E tests..."
	# the target assumes the MCLI-* environment variables are exported
	$(E2E_CMD) -v -p 1 -parallel 1 -tags="$(E2E_TAGS)" ./e2e...

.PHONY: install
install: ## Install a binary in $GOPATH/bin
	go install -ldflags "${LINKER_FLAGS}"

.PHONY: gen-notices
gen-notices: ## Generate 3rd party notices
	@echo "==> Generating 3rd party notices"
	@chmod -R 777 ./third_party_notices
	@rm -Rf third_party_notices
	go-licenses save "github.com/mongodb/mongocli" --save_path=third_party_notices

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
