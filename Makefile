# vim: set softtabstop=2 shiftwidth=2:
SHELL = bash

# ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
# GOPATH := ${ROOT}:${GOPATH}
export GOPATH := ${PWD}
CURDIR = ${PWD}

all: install start

install:
	@echo "Build" 
	$(shell rm ./bin/index)
	$(shell export GOPATH=${CURDIR}; go build -o ./bin/index ./src/index.go)

start:
	@echo "Start" 
	./bin/index

test:
	@echo ${GOPATH}
	@go test ./src/closest
