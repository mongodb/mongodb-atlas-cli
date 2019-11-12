SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on

# gofmt and goimports all go files
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
.PHONY: fmt

# Clean go.mod
go-mod-tidy:
	@go mod tidy -v
	@git diff HEAD
	@git diff-index --quiet HEAD
.PHONY: go-mod-tidy

build:
	go build
.PHONY: build

install:
	go install
.PHONY: install

.DEFAULT_GOAL := build
