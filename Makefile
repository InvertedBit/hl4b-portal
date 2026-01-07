.PHONY: help build run dev clean install test css watch-css

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install all dependencies
	go mod download
	npm install

build: css ## Build the application
	go build -o hl4b-portal main.go

run: build ## Build and run the application
	./hl4b-portal

dev: ## Run the application with live reload (requires air)
	air

css: ## Build Tailwind CSS
	npm run build:css

watch-css: ## Watch and rebuild Tailwind CSS on changes
	npm run watch:css

clean: ## Clean build artifacts
	rm -f hl4b-portal
	rm -rf uploads/*
	rm -f static/css/output.css

test: ## Run tests
	go test -v ./...

fmt: ## Format code
	go fmt ./...

lint: ## Run linter (requires golangci-lint)
	golangci-lint run

db-init: ## Initialize PostgreSQL database with auth.users table
	@echo "Run this SQL script on your PostgreSQL database:"
	@echo "psql -U your_user -d your_database -f init.sql"

.DEFAULT_GOAL := help
