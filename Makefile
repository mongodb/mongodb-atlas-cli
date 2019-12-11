SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on

setup:
	@echo "==> Installing dependencies..."
    curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0
.PHONY: setup

# gofmt and goimports all go files
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

lint:
	@echo "==> Linting all packages..."
	golangci-lint run $(SOURCE_FILES) -E goimports -E golint -E misspell -E unconvert -E maligned
.PHONY: lint

build:
	go build
.PHONY: build

install:
	go install
.PHONY: install

.DEFAULT_GOAL := build
