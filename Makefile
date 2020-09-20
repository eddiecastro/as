SHELL := /bin/bash
PREFIX?=$(shell pwd)
PKG=$(shell pwd)

WEBPACK_PORT=5000

.PHONY: frontend
frontend:
	npm start

.PHONY: run-backend
run-backend:
	go run ./backend/cmd/backend
