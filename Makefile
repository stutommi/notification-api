BIN				= notification-api
SRC_DIR			= src
BIN_OUTPUT_PATH	= bin/${BIN}
ENV_FILE		= .env

SHELL := /bin/bash

all: run

build:
	@go build -o ${BIN_OUTPUT_PATH} ${SRC_DIR}/*.go

run: build
	@${BIN_OUTPUT_PATH}

dev: build
	@source ${ENV_FILE} && ${BIN_OUTPUT_PATH}

test:
	go test ${SRC_DIR}/*