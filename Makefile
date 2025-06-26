# Makefile for building Go binary
.PHONY: build clean run all

# Binary name
BINARY_NAME=public_endpoints

# Go build command
build:
	go build -o $(BINARY_NAME)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run the application
run:
	go run .

# Build and run
all: build run