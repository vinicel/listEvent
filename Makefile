DOCKER = docker
DOCKER_COMPOSE = docker-compose

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY : up
up: ## Up docker
	@$(DOCKER_COMPOSE) -f ./docker-compose.yml up

.PHONY : down
down: ## Down docker
	@$(DOCKER_COMPOSE) -f ./docker-compose.yml down