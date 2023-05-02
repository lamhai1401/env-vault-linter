.PHONY: pull build push install share

ifeq ($(ENV),)
ENV := development
endif

ifeq ($(ENV),develop)
ENV := development
endif

ENV_FILE :=  .env.$(ENV)

install:
	npm i -g dotenv-vault

build:
	@echo ${ENV}
	npx dotenv-vault build

pull:
	npx dotenv-vault pull $(ENV) $(ENV_FILE)

push:
	npx dotenv-vault push  $(ENV) $(ENV_FILE)

share-api:
	@echo "--> $(ENV) environment.."
	cat ./$(ENV_FILE) > ../../packages/api/$(ENV_FILE) || true

share-webapp:
	@echo "--> $(ENV) environment.."
	cat ./$(ENV_FILE) > ../../packages/webapp/assets/.env || true

share-migrate:
	@echo "--> $(ENV) environment.."
	cat ./$(ENV_FILE) >  ../../packages/migrate/$(ENV_FILE) || true

# env-quotes-format:
# 	sh ./formatenv.sh $(ENV)

share-infra:
	@echo "--> $(ENV) environment.."
	cat ./$(ENV_FILE) >  ../terraform/terraform.tfvars || true

share: share-api 
# share-webapp share-migrate share-infra

dot-me:
	npx dotenv-vault login
