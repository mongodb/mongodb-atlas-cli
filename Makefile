SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on

setup:
	@echo "==> Installing dependencies..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.22.0
.PHONY: setup

# GIT hooks
link-git-hooks:
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;
.PHONY: link-git-hooks

# gofmt and goimports all go files
fmt:
	@echo "==> Formatting all files..."
	find . -name '*.go' -not -wholename './mocks/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

lint:
	@echo "==> Linting all packages..."
	golangci-lint run $(SOURCE_FILES) -E goimports -E golint -E misspell -E unconvert -E maligned --skip-dirs ^mocks/
.PHONY: lint

test:
	@echo "==> Running tests..."
	go test $(SOURCE_FILES) -timeout=30s -parallel=4
.PHONY: test

check: test lint
.PHONY: check

build:
	go build
.PHONY: build

install:
	go install
.PHONY: install

.DEFAULT_GOAL := build
