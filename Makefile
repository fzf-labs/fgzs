SHELL := /bin/bash
BASEDIR = $(shell pwd)

export GOPATH := $(shell go env GOPATH)
export GOPROXY := https://goproxy.cn/,direct

.PHONY: init
# make init
init:
	@go install github.com/zeromicro/go-zero/tools/goctl@latest
	@goctl env check -i -f --verbose
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

.PHONY: fmt
# make fmt
fmt:
	@gofmt -s -w .

.PHONY: vet
# make vet
vet:
	@go vet ./...

.PHONY: ci-lint
# make ci-lint
ci-lint:
	@golangci-lint run ./...

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

