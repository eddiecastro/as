SHELL := /bin/bash
PREFIX?=$(shell pwd)
PKG=$(shell pwd)

.PHONY: run-backend
run-backend:
	go run ./backend/cmd/backend

.PHONY: run-frontend
run-frontend:
	cd $(PKG)/frontend && go build -v -o ./frontend ./ && ./frontend

.PHONY: create_docker
create_docker:  ## Build docker images for deploy
	docker build -t localhost:5000/sal-backend:latest -f Dockerfile-backend . && \
	docker build -t localhost:5000/sal-frontend:latest -f Dockerfile-frontend .

.PHONY: publish_docker
publish_docker: ## Publish docker images
	docker push localhost:5000/sal-backend:latest && \
	docker push localhost:5000/sal-frontend:latest
