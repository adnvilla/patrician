SHELL := /bin/bash

default: build

build: commit_hash install

install: 
	GOBIN=`pwd`/bin go install -v

commit_hash: 
	./generate_commit_hash_file.rb

run: build
	./bin/patrician