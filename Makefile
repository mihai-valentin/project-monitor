.DEFAULT_GOAL := build-and-up

help:
	@echo "Help ............ make help"
	@echo "Build and Run ... make build-and-up"
	@echo "Build ........... make build"
	@echo "Run ............. make up"

build-and-up: build up

build:
	@docker-compose build

up:
	@docker-compose up -d
