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
	cd $(PKG)/backend && \
	docker build -t localhost:5000/backend:latest -f Dockerfile . && \
	cd $(PKG)/frontend && \
	docker build -t localhost:5000/frontend:latest -f Dockerfile .