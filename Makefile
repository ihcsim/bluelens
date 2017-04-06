.PHONY: all server client vendor test

PACKAGE_ROOT = github.com/ihcsim/bluelens
PACKAGE_DESIGN = ${PACKAGE_ROOT}/design
PACKAGE_CLIENT = ${PACKAGE_ROOT}/cmd/blue

export PACKAGE_SERVER = ${PACKAGE_ROOT}/cmd/blued
export PACKAGE_CLI = ${PACKAGE_CLIENT}/tool/blue

CLIENT_DIR = cmd/blue
SERVER_DIR = cmd/blued
SERVER_HOSTNAME ?= localhost
SERVER_SCHEME ?= https

GOAGEN_VERSION = v1.1.0
SHELL := /bin/bash
GLIDE := $(shell command -v glide 2> /dev/null)
ifndef GLIDE
$(error "Please install glide. Installation instruction can be found at https://github.com/Masterminds/glide#install")
endif

all: vendor test codegen build
codegen: server/codegen client/codegen
build: server/build client/build
server: server/codegen server/build
client: client/codegen client/build

server/codegen:
	goagen main ${CODEGEN_MAIN_OPTS} -d ${PACKAGE_DESIGN} -o ${SERVER_DIR}
	goagen app -d ${PACKAGE_DESIGN} -o ${SERVER_DIR}
	goagen swagger -d ${PACKAGE_DESIGN} -o ${SERVER_DIR}

server/build:
	go build -v -o blued ${PACKAGE_SERVER}

client/codegen:
	goagen client --tool blue -d ${PACKAGE_DESIGN} -o ${CLIENT_DIR}
	goagen js --scheme=${SERVER_SCHEME} --host=${SERVER_HOSTNAME} -d ${PACKAGE_DESIGN} -o ${CLIENT_DIR}

client/build:
	go build -v -o blue ${PACKAGE_CLI}

test:
	go test -v -cover -race `glide novendor`

vendor:
	glide install

goagen: vendor
	go install ${PACKAGE_ROOT}/vendor/github.com/goadesign/goa/goagen

tls:
	mkdir -p tls
	openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout tls/localhost.key -out tls/localhost.crt -subj "/CN=localhost" -days 365

aci: tls
	acbuild.sh
