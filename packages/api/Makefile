.PHONY: install format test clean build run

# Set default env as develop.
ifeq ($(ENV),)
ENV := development
endif

ifeq ($(ENV),develop)
ENV := development
endif

# Set current Git revision unless given explicitly.
ifeq ($(GIT_REVISION),)
GIT_REVISION := $(shell git rev-parse --short HEAD)
endif

-include ./.env.${ENV}

BUILD := ./build

GRAPH_DIR := ./graph
GRAPH_SCHEMA := $(GRAPH_DIR)/schema.graphqls

DAO_DIR := ./database/dao

GORUN := GIT_REVISION=$(GIT_REVISION) go run

# DB_GEN := gentool -db postgres -dsn "host=${SECRET_POSTGRES_HOSTNAME} port=${SECRET_POSTGRES_PORT} \
# 	user=${SECRET_POSTGRES_USER}  password=${SECRET_POSTGRES_PASSWORD} dbname=${SECRET_POSTGRES_DATABASE} sslmode=disable" \
# 	-fieldNullable -fieldWithIndexTag -fieldWithTypeTag -outPath ${DAO_DIR} -fieldSignable -withUnitTest -modelPkgName entities

TEST := GIT_REVISION=$(GIT_REVISION) go test -cover -coverprofile coverage.log

TEST_TOOL := go tool cover -html=coverage.log

DB_GEN := cd script/dbgen/ && go run .

run:
	air

format:
	go fmt ./...

dbgen:
	@echo "Generating entity models, DAO from database"
	$(DB_GEN)
	@echo "Generate completed"

schemagen:
	@echo "Copying graphqls schema.."
	@mkdir -p $(GRAPH_DIR)
	@\cp -rf ../schema/build/schema.graphqls ./$(GRAPH_SCHEMA)
	@echo "Success!"

apigen: schemagen
	@echo "Generating graphqls resolver.."
	go run github.com/99designs/gqlgen generate
	@echo "Success!"

test:
	@echo "TODO"

install:
	go mod download

init:
	go run github.com/99designs/gqlgen init
