# Variables
COMPOSEFILE := ./srcs/docker-compose.yml
COMPOSE := docker compose -f $(COMPOSEFILE)
COLOR := \033[1;33m
RESET := \033[0m

# Disable command echoing
.SILENT:

# Default target
.DEFAULT_GOAL := build

# Phony targets
.PHONY: build up down stop restart clean prune re status

# Build and start containers
build:
	echo -e "$(COLOR)BUILD$(RESET)"
	$(COMPOSE) up -d --build
	$(COMPOSE) up

# Start containers
up:
	echo -e "$(COLOR)UP$(RESET)"
	$(COMPOSE) up

# Stop and remove containers, networks, and volumes
down:
	echo -e "$(COLOR)DOWN$(RESET)"
	$(COMPOSE) down

# Stop running containers
stop:
	echo -e "$(COLOR)STOP$(RESET)"
	$(COMPOSE) stop

# Restart containers
restart:
	echo -e "$(COLOR)RESTART$(RESET)"
	$(COMPOSE) restart

# Clean up containers, volumes, and images
clean:
	echo -e "$(COLOR)CLEAN$(RESET)"
	$(COMPOSE) down -v --rmi local

# System-wide cleanup
prune:
	echo -e "$(COLOR)PRUNE$(RESET)"
	$(COMPOSE) down -v
	docker system prune -af
	docker volume prune -f

# Restart services (down + build)
re: down build

# Show status of Docker resources
status:
	echo -e "$(COLOR)Containers$(RESET)"
	docker ps -a
	echo -e "$(COLOR)Images$(RESET)"
	docker images
	echo -e "$(COLOR)Volumes$(RESET)"
	docker volume ls
	echo -e "$(COLOR)Networks$(RESET)"
	docker network ls
