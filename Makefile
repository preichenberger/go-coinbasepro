SHELL := /bin/bash

all: run test

.PHONY: run
run:
	. .env && go run cmd/go-gdax/main.go

.PHONY: run_race
run_race:
	. .env && go run -race cmd/go-gdax/main.go

.PHONY: test
test:
	. .env &&  go test -v -race ./...
