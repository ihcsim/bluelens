.PHONY: all server client vendor test

PACKAGE_ROOT = github.com/ihcsim/bluelens
PACKAGE_DESIGN = ${PACKAGE_ROOT}/design
PACKAGE_SERVER = ${PACKAGE_ROOT}/server
PACKAGE_CLIENT = ${PACKAGE_ROOT}/client
PACKAGE_CLI = ${PACKAGE_CLIENT}/tool/blue

CLIENT_DIR = client
SERVER_DIR = server
SERVER_HOSTNAME ?= localhost
SERVER_SCHEME ?= http

SHELL := /bin/bash
GLIDE := $(shell command -v glide 2> /dev/null)
ifndef GLIDE
$(error "Please install glide. Installation instruction can be found at https://github.com/Masterminds/glide#install")
endif

all: vendor test server client
server: server/codegen server/build
client: client/codegen client/build

server/codegen:
	goagen main -d ${PACKAGE_DESIGN} -o ${SERVER_DIR}
	goagen app -d ${PACKAGE_DESIGN} -o ${SERVER_DIR}
	goagen swagger -d ${PACKAGE_DESIGN} -o ${SERVER_DIR}

server/build:
	go build -v -o bluelens ${PACKAGE_SERVER}

client/codegen:
	goagen client --tool blue -d ${PACKAGE_DESIGN} -o ${CLIENT_DIR}
	goagen js --scheme=${SERVER_SCHEME} --host=${SERVER_HOSTNAME} -d ${PACKAGE_DESIGN} -o ${CLIENT_DIR}

client/build:
	go build -v -o blue ${PACKAGE_CLI}

test:
	go test -v -cover -race `glide novendor`

vendor:
	glide install
