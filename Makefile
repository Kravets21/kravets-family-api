SHELL = /bin/bash

APP_CONTAINER_NAME := kravets_family_api
DB_CONTAINER_NAME := kravets_family_database
docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

build: ## Build all containers and start (interact) for development
	@echo "build kravets family api service..."
	$(docker_compose_bin) build --no-cache --parallel --force-rm
	$(docker_compose_bin) up --remove-orphans --force-recreate

up: ## Start all containers (no interact) for development
	@echo "build and run..."
	$(docker_compose_bin) up --no-recreate -d

debug: ## Start all containers (interact) for development
	@echo "run with logs..."
	$(docker_compose_bin) up --no-recreate

clear: down
	@echo "rm app images..."
	$(docker_bin) images -a | grep "cms_*" | awk '{print $3}' | xargs docker rmi --force
	$(docker_bin) rm ${APP_CONTAINER_NAME}

rebuild: ## Build all containers and start (interact) for development
	@echo "rebuild and run..."
	$(docker_bin) image prune -a --force
	$(docker_compose_bin) up --build --remove-orphans --force-recreate

down: ## Stop all started for development containers
	@echo "down..."
	$(docker_compose_bin) down

restart: up ## Restart all started for development containers
	@echo "restart..."
	$(docker_compose_bin) restart

shell: up ## Start shell into application container
	$(docker_compose_bin) exec "$(APP_CONTAINER_NAME)" /bin/sh
