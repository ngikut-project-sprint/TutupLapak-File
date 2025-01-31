# Variables
APP_NAME := TutupLapak-File
BUILD_DIR := bin
MAIN_FILE := cmd/main.go

# Supported Platforms
OS ?= linux
ARCH ?= arm64
GO_ENV_VARS := GOOS=$(OS) GOARCH=$(ARCH)

# Default target
all: build

.PHONY: build
build:
	@echo "Building for $(OS)/$(ARCH)..."
	@mkdir -p $(BUILD_DIR)
	$(GO_ENV_VARS) go build -o $(BUILD_DIR)/$(APP_NAME)-$(OS)-$(ARCH) $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)-$(OS)-$(ARCH)"

.PHONY: build-all
build-all: ARCH_LIST := amd64 arm arm64 386
build-all:
	@for arch in $(ARCH_LIST); do \
		GOOS=$(OS) GOARCH=$$arch go build -o $(BUILD_DIR)/$(APP_NAME)-$(OS)-$$arch $(MAIN_FILE); \
		echo "Built for $(OS)/$$arch"; \
		done

.PHONY: run
run:
	@echo "Running the application"
	go run $(MAIN_FILE)

.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(BUILD_DIR)

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

.PHONY: mod
mod:
	@echo "Downloading and tidying dependencies..."
	go mod tidy

.PHONY: lint
lint:
	@echo "Running lint..."
	golangci-lint run

.PHONY: mocks
mocks:
	@echo "Generate test mocks..."
	mockery --all

.PHONY: proto
proto:
	@echo "Generate protocol buffer..."
	protoc -I=cmd/api/grpc \
  --go_out=cmd/api/grpc --go_opt=paths=source_relative \
  --go-grpc_out=cmd/api/grpc --go-grpc_opt=paths=source_relative \
  cmd/api/grpc/file_service.proto

.PHONY: help
help:
	@echo "Available commands:"
	@echo " make build      - Build for the current architecture (default: arm64)"
	@echo " make build-all  - Build application for all architecture (amd64, arm, arm64, 386)"
	@echo " make run        - Run application"
	@echo " make test       - Run all test case"
	@echo " make fmt        - Format code"
	@echo " make mod        - Fix dependencies"
	@echo " make lint       - Lint code"
	@echo " make clean      - Clean all build artifacts"
	@echo " make mocks      - Generate test mocks"
	@echo " make proto      - Generate protocol buffer"
	@echo	""
	@echo	"Environment variables:"
	@echo " OS              - Build target operating system (default: linux)"
	@echo " ARCH            - Build target processor architecture (default: arm64)"
