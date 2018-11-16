PROJECT_NAME := go-rest-api
CREATE_DB_FILE := scripts/create-db.sql

.DEFAULT_GOAL := run

#!make
include .env
export

.PHONY: db build run clean

db:
	sqlite3 ${DATABASE_PATH} < ${CREATE_DB_FILE}

build:
	go build -o ${PROJECT_NAME} main.go

run: build
	./${PROJECT_NAME}

clean:
	rm -f ${PROJECT_NAME} ${DATABASE_PATH}
