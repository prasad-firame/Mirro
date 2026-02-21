# Makefile for Mirro

# Default target
all: build

# Build the CLI
build:
    go build ./cmd/mirro

# Run tests
test:
    go test ./... -v

# Run the proxy locally with example config
run:
    ./mirro serve --config examples/simple/config.yaml

# Clean build artifacts
clean:
    rm -f mirro mirro.exe