# Makefile for API Mobile Dashboard

.PHONY: help build up down dev logs clean test check-docker

# Check if Docker is installed
check-docker:
	@which docker > /dev/null || (echo "Docker is not installed. Please install Docker Desktop from https://www.docker.com/products/docker-desktop" && exit 1)
	@docker --version

# Default target
help:
	@echo "Available commands:"
	@echo "  check-docker - Check if Docker is installed"
	@echo "  build        - Build the Docker images"
	@echo "  up           - Start the production services"
	@echo "  down         - Stop all services"
	@echo "  dev          - Start development environment with hot reload"
	@echo "  logs         - Show logs from all services"
	@echo "  logs-api     - Show logs from API service only"
	@echo "  clean        - Remove all containers, images, and volumes"
	@echo "  test         - Run tests"
	@echo "  shell        - Open shell in API container"
	@echo "  run          - Run locally without Docker"
	@echo "  install      - Install Go dependencies"

# Build Docker images
build: check-docker
	docker compose build

# Start production environment
up: check-docker
	docker compose up -d

# Stop all services
down: check-docker
	docker compose down

# Start development environment with hot reload
dev: check-docker
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up --build

# Show logs
logs: check-docker
	docker compose logs -f

# Show API logs only
logs-api: check-docker
	docker compose logs -f api

# Clean up everything
clean: check-docker
	docker compose down -v --rmi all --remove-orphans
	docker system prune -f

# Run tests
test:
	go test ./...

# Open shell in API container
shell: check-docker
	docker compose exec api sh

# Install dependencies
install:
	go mod download
	go mod tidy

# Run locally without Docker
run:
	go run cmd/api/main.go

# Build binary
build-binary:
	go build -o api-mobile-dashboard ./cmd/api/main.go