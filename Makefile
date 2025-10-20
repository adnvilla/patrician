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
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-integration:
	go test -v ./integration/...

test-domain:
	go test -v ./internal/domain/...

test-use-cases:
	go test -v ./internal/application/use_cases/...

test-handlers:
	go test -v ./internal/interfaces/handlers/...

test-repositories:
	go test -v ./internal/infrastructure/data/postgresql/...

.PHONY: build install commit_hash run test test-coverage test-cover test-integration test-domain test-use-cases test-handlers test-repositories