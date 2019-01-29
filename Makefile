#!/usr/bin/env make

ifeq ($(GIT_ROOT),)


GIT_ROOT:=$(shell git rev-parse --show-toplevel)
endif

.PHONY: all clean format lint vet bindata build test docker-deps reap dist

all: lint vet test build

format:
	${GIT_ROOT}/make/format

lint:
	${GIT_ROOT}/make/lint

vet:
	${GIT_ROOT}/make/vet

build:
	${GIT_ROOT}/make/build

test:
	${GIT_ROOT}/make/test
