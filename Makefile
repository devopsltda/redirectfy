# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@templ generate
	@npm run build
	@go build -o ./bin/main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f ./bin/main

# Update docs
docs:
	@if command -v swag > /dev/null; then \
			swag init --generalInfo routes.go --parseDependency --parseInternal --dir internal/server; \
	else \
	    read -p "Go's 'swag' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/swaggo/swag/cmd/swag@latest; \
					swag init --generalInfo routes.go --parseDependency --parseInternal --dir internal/server; \
	    else \
	        echo "You chose not to install swag. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Live Reload
watch:
	@if command -v air > /dev/null; then \
			npm run watch & air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean docs
