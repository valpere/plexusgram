.PHONY: help build run test clean fmt vet lint tidy deps install

# Variables
BINARY_NAME=plexusgram
MAIN_PATH=main.go
COVERAGE_FILE=coverage.out

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the application binary
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run: ## Run the application
	go run $(MAIN_PATH)

test: ## Run all tests
	go test ./...

test-verbose: ## Run tests with verbose output
	go test -v ./...

test-coverage: ## Run tests with coverage report
	go test -cover ./...

test-coverage-html: ## Generate HTML coverage report
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE)

fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run golangci-lint (requires golangci-lint installed)
	golangci-lint run

tidy: ## Tidy and verify dependencies
	go mod tidy
	go mod verify

deps: ## Download dependencies
	go mod download

install: ## Install the binary to GOPATH/bin
	go install

clean: ## Remove build artifacts and coverage files
	rm -f $(BINARY_NAME)
	rm -f $(COVERAGE_FILE)
	go clean

check: fmt vet ## Run formatting and vetting

all: clean deps build ## Clean, download deps, and build

.DEFAULT_GOAL := help
