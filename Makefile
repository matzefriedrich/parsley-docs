.PHONY: init update-docs build start docker-build docker-run clean

BIN_DIR := $(shell pwd)/bin
GOMARKDOC := $(BIN_DIR)/gomarkdoc

# Default target
all: build

init:
	@echo "Setting up local bin directory..."
	mkdir -p $(BIN_DIR)
	@echo "Installing gomarkdoc..."
	GOBIN=$(BIN_DIR) go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	@echo "Tool 'gomarkdoc' installed to $(BIN_DIR)"
	@echo ""
	@echo "Note: 'retype' should be installed globally if not already present:"
	@echo "  dotnet tool install retypeapp --global"

update-docs:
	@if [ ! -d "../parsley" ]; then \
		echo "Warning: ../parsley directory not found. update-package-docs.sh may fail."; \
	fi
	PATH=$(BIN_DIR):$$PATH ./update-package-docs.sh

build:
	retype build

start:
	retype start

docker-build:
	docker build --rm -t parsley-docs -f Dockerfile .

docker-run:
	@echo "Documentation will be available at http://localhost:27821"
	docker run --rm -p 27821:80 parsley-docs

clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	rm -rf .retype
