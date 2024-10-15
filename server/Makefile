# Makefile for Go project

# Define the default target
.PHONY: all test clean

# Define the test target
test:
	@echo "Running tests with GO_ENV=test..."
	GO_ENV=test go test ./...
# Clean target to remove any generated files or binaries
clean:
	@echo "Cleaning up..."
	go clean

# Default target to run when `make` is called
all: test
