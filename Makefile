SHELL := /bin/bash

REPOS = $(shell go list -f '{{join .Imports "\n"}}{{"\n"}}{{join .TestImports "\n"}}' ./... | sort | uniq | grep -v golang-samples)

default: install

install:
	go get -u -v ${REPOS}

test:
	go vet ./...
	go test ./...
