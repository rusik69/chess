# Chess Game Makefile
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Binary names
BINARY_NAME=chess-game
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME).exe
BINARY_MAC=$(BINARY_NAME)_mac

# Build directory
BUILD_DIR=bin

# Default target
.PHONY: all
all: clean build

# Build the binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -func=coverage.out

# Generate HTML coverage report
.PHONY: test-cover
test-cover: test-coverage
	@echo "Generating HTML coverage report..."
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run the application
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	$(GOVET) ./...

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Check code quality (fmt, vet, lint)
.PHONY: check
check: fmt vet lint

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Development workflow
.PHONY: dev
dev: deps check test

# Cross compilation
.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v ./

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_WINDOWS) -v ./

.PHONY: build-mac
build-mac:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_MAC) -v ./

# Show help
.PHONY: help
help:
	@echo "Chess Game Makefile Commands:"
	@echo ""
	@echo "Building:"
	@echo "  build         Build the chess game binary"
	@echo "  build-linux   Build for Linux (amd64)"
	@echo "  build-windows Build for Windows (amd64)"
	@echo "  build-mac     Build for macOS (amd64)"
	@echo "  build-all     Build for all platforms"
	@echo "  install       Install binary to GOPATH/bin"
	@echo ""
	@echo "Running:"
	@echo "  run           Run the chess game"
	@echo "  dev           Quick dev cycle (fmt, vet, test, run)"
	@echo ""
	@echo "Testing:"
	@echo "  test          Run all tests"
	@echo "  test-coverage Run tests with HTML coverage report"
	@echo "  test-cover    Run tests with terminal coverage"
	@echo ""
	@echo "Code Quality:"
	@echo "  fmt           Format all Go files"
	@echo "  vet           Run go vet"
	@echo "  lint          Run golangci-lint (if installed)"
	@echo "  check         Run fmt, vet, and test"
	@echo ""
	@echo "Maintenance:"
	@echo "  clean         Remove build artifacts"
	@echo "  deps          Download and tidy dependencies"
	@echo ""
	@echo "Other:"
	@echo "  all           Run test and build (default)"
	@echo "  help          Show this help message" 