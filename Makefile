.DEFAULT_GOAL := help
PROJECT_BIN = $(shell pwd)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)
GOOS = linux
GOARCH = amd64
CGO_ENABLED = 0
VERS = $(shell git describe --tags --abbrev=0)
LDFLAGS = "-w -s -X main.vers="$(VERS)
APP := $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))
TARGET = cmd/telecho/main.go

.PHONY: build \
		run \
		.install-linter \
		lint \
		.install-nil \
		nil-check \
		help

run: build
	echo "123\n123\n123" | $(APP)
	$(APP) --help

build: ## Build release
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=$(LDFLAGS) -o $(PROJECT_BIN)/$(APP) $(TARGET)

.install-linter: ## Install linter
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.54.2

lint: .install-linter ## Run linter
	golangci-lint run ./...

.install-nil: ## Install nil check
	[ -f $(PROJECT_BIN)/nilaway ] || go install go.uber.org/nilaway/cmd/nilaway@latest && cp $(GOPATH)/bin/nilaway $(PROJECT_BIN)

nil-check: .install-nil ## Run nil check linter
	nilaway ./...

help:
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
