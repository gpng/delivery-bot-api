# Docker parameters
DOCKERCMD=docker
DOCKERCOMPOSECMD=docker-compose

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

# App parameters
CONTAINER_TAG=delivery-bot
DOCKERFILE=Dockerfile.dev
MAIN_FOLDER=cmd/api
MAIN_PATH=$(MAIN_FOLDER)/main.go
DEPLOY_FOLDER=deploy
DEPLOY_BINARY_NAME=$(DEPLOY_FOLDER)/application
ZIP_PATH=$(DEPLOY_FOLDER)/deploy-$(shell date +'%Y%m%d-%H%M%S').zip

default:
	@echo "=============building API============="
	$(DOCKERCMD) build -f $(DOCKERFILE) -t $(CONTAINER_TAG) .

up: default
	@echo "=============starting API locally============="
	$(DOCKERCOMPOSECMD) up -d

logs:
	$(DOCKERCOMPOSECMD) logs -f

dev: up logs

run:
	go build -o bin/application $(MAIN_PATH) && ./bin/application -docs

down:
	$(DOCKERCOMPOSECMD) down

test:
	godotenv -f .test.env $(GOTEST) -cover ./...

clean: down
	@echo "=============cleaning up============="
	$(DOCKERCMD) system prune -f
	$(DOCKERCMD) volume prune -f

run-prod:
	$(DOCKERCMD) build -t chanced-api-eb .
	docker run -p 4000:5000 chanced-api-eb

generate-docs:
	swag init -g $(MAIN_PATH)