# TODO: Change the APP, VERSION, GITHUB_OWNER, and GITHUB_DOMAIN variables to match your application
APP=go-template
VERSION=0.1.0
GITHUB_OWNER=ginolatorilla
GITHUB_DOMAIN=github.com

COMMIT_HASH=$(shell git rev-parse HEAD)
PACKAGE=$(GITHUB_DOMAIN)/$(GITHUB_OWNER)/$(APP)

BUILD_FLAGS=-v -buildvcs 
LD_FLAGS=-ldflags="-X '$(PACKAGE)/version.AppName=$(APP)' -X '$(PACKAGE)/version.Version=$(VERSION)' -X '$(PACKAGE)/version.CommitHash=$(COMMIT_HASH)'"

.PHONY: all
all: test tidy build

.PHONY: test
test:
	@echo "🌡  Running tests..."
	@go test -race $(BUILD_FLAGS) $(LD_FLAGS) ./...

.PHONY: test/cover
test/cover:
	@echo "🌡️  Running tests..."
	@go test -coverprofile=/tmp/coverage.out -race $(BUILD_FLAGS) $(LD_FLAGS) ./...
	@go tool cover -html=/tmp/coverage.out

.PHONY: tidy
tidy:
	@echo "🧹 Tidying up package dependencies..."
	@go mod tidy

.PHONY: build
build:
	@echo "🏗️  Building the application..."
	@go build $(BUILD_FLAGS) $(LD_FLAGS) -o bin/$(APP) $(PACKAGE) 

.PHONY: clean
clean:
	go clean
	rm -rf bin/*

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  all        - Run test, tidy, and build (default)"
	@echo "  help       - Show this help message"
	@echo "  test       - Run tests"
	@echo "  test/cover - Run tests with coverage"
	@echo "  tidy       - Sort out package dependencies"
	@echo "  build      - Build the application"
	@echo "  clean      - Clean up the build artifacts"