.PHONY: build run dev clean test deps setup

# Variables
BINARY_NAME=waba-bot
GO_FILES=$(shell find . -name "*.go" -type f)

# Default target
all: build

# Install dependencies
deps:
	go mod tidy
	go mod download

# Build the application
build: deps
	go build -o $(BINARY_NAME) .

# Run the application
run: build
	./$(BINARY_NAME)

# Development mode (auto-reload with air)
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Installing air for auto-reload..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# Setup project
setup:
	./setup.sh

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

# Docker build
docker-build:
	docker build -t $(BINARY_NAME) .

# Docker run
docker-run:
	docker run -p 3000:3000 --env-file .env $(BINARY_NAME)

# Show help
help:
	@echo "Available targets:"
	@echo "  deps         - Install Go dependencies"
	@echo "  build        - Build the application"
	@echo "  run          - Build and run the application"
	@echo "  dev          - Run in development mode with auto-reload"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  setup        - Setup project with dependencies"
	@echo "  fmt          - Format Go code"
	@echo "  lint         - Run linter"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  help         - Show this help"
