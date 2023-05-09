# Variables
APP_NAME = myapp
GO_FILES = $(wildcard *.go)
GO_CMD = go
GO_BUILD = $(GO_CMD) build
GO_TEST = $(GO_CMD) test

# Default target: build
all: build

# Build target: compile the Go source files into the final executable
build: $(APP_NAME)

# Compile the Go source files
$(APP_NAME): $(GO_FILES)
	$(GO_BUILD) -o $(APP_NAME) $(GO_FILES)

# Test target: run the tests for the project
test:
	$(GO_TEST) -v ./... -cover

# Declare phony targets
.PHONY: all build test
