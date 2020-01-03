SOURCE_FILES?=./...
BINARY_NAME=mcli

DESTINATION=./bin/${BINARY_NAME}

LINKER_FLAGS=-X github.com/10gen/mcli/internal/version.Version=${VERSION}

export PATH := ./bin:$(PATH)
export GO111MODULE := on

.PHONY: setup
setup:
	@echo "==> Installing dependencies..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.22.2
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

# GIT hooks
.PHONY: link-git-hooks
link-git-hooks:
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

# gofmt and goimports all go files
.PHONY: fmt
fmt:
	@echo "==> Formatting all files..."
	find . -name '*.go' -not -wholename './mocks/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

.PHONY: lint
lint:
	@echo "==> Linting all packages..."
	golangci-lint run $(SOURCE_FILES) -E goimports -E golint -E misspell -E unconvert -E maligned --skip-dirs ^mocks/

.PHONY: test
test:
	@echo "==> Running tests..."
	go test $(SOURCE_FILES) -timeout=30s -parallel=4

.PHONY: check
check: test lint

.PHONY: gen-mocks
gen-mocks:
	mockgen -source=internal/config/profile.go -destination=mocks/mock_config.go -package=mocks
	mockgen -source=internal/store/clusters.go -destination=mocks/mock_clusters.go -package=mocks
	mockgen -source=internal/store/database_users.go -destination=mocks/mock_database_users.go -package=mocks
	mockgen -source=internal/store/project_ip_whitelist.go -destination=mocks/mock_project_ip_whitelist.go -package=mocks
	mockgen -source=internal/store/projects.go -destination=mocks/mock_projects.go -package=mocks

.PHONY: build
build: compile-local

.PHONY: compile-local
compile-local:
	go build -ldflags "${LINKER_FLAGS}" -o ${DESTINATION}

.PHONY: release
release: goreleaser

.DEFAULT_GOAL := build
