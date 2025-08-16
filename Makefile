SHELL := /bin/bash

default: build

build: commit_hash install

install: 
	GOBIN=`pwd`/bin go install -v

commit_hash: 
	./generate_commit_hash_file.rb

run: build
	./bin/patrician

test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: build install commit_hash run test test-verbose test-coverage

test:
	go test -v ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-integration:
	go test -v ./integration/...

test-domain:
	go test -v ./src/domain/...

test-use-cases:
	go test -v ./src/application/use_cases/...

test-handlers:
	go test -v ./src/interfaces/handlers/...

test-repositories:
	go test -v ./src/infrastructure/data/postgresql/...