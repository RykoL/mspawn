# Variables
BINARY_NAME=mspawn
BUILD_DIR=bin
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")

# Default target
.DEFAULT_GOAL := help

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## build: Build the binary application
.PHONY: build
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

## clean: Remove build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf $(BUILD_DIR)

## pre-commit: Executes all quality gates
.PHONY: pre-commit
pre-commit:
	$(MAKE) vet
	$(MAKE) lint
	$(MAKE) test/cover
	$(MAKE) build

## test: Run tests
.PHONY: test
test:
	@echo "Testing..."
	go test -v ./...

## test/cover: Run tests with coverage output
.PHONY: test/cover
test/cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

## lint: Lint the code
.PHONY: lint
lint:
	golangci-lint run


## fmt: Format the code
.PHONY: fmt
fmt:
	@echo "Formatting..."
	go fmt ./...

## vet: Analyze code for static errors
.PHONY: vet
vet:
	@echo "Vetting..."
	go vet ./...

## deps: Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

## help: Show this help message
.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
