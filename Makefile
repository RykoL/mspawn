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
