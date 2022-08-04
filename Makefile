.PHONY: help build
S=silverton
REGISTRY:=us-east1-docker.pkg.dev/silverton-io/docker
VERSION:=$(shell cat .VERSION)
BINARY_NAME:=tracked-ui

build:
	go build -ldflags="-X main.VERSION=$(VERSION)" -o $(BINARY_NAME) ./cmd/*.go

build-docker:
	docker build -f deploy/Dockerfile -t tracked-ui:$(VERSION) .
