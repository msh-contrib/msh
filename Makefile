.PHONY: build

build:
	@go install ./cmd/msh

build-graph:
	@go install ./graph
