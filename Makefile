PATH  := $(PATH):$(PWD)
SHELL := env PATH=$(PATH) /bin/bash

default: run
.PHONY: fmt build template run

fmt:
	go fmt ./...

build: fmt
	go install github.com/jinmatt/twtrgo/cmd/twtrgo

template: build
	hero -source="./http/template/src" -dest="./http/template"

run: build
	$(GOPATH)/bin/twtrgo
