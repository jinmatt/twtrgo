PATH  := $(PATH):$(PWD)
SHELL := env PATH=$(PATH) /bin/bash

default: run
.PHONY: fmt test build template run

fmt:
	go fmt ./...

test: fmt
	go test -v ./test

build: fmt
	go install github.com/jinmatt/twtrgo/cmd/twtrgo

template: build
	hero -source="./http/template/src" -dest="./http/template"

run: build
	$(GOPATH)/bin/twtrgo
