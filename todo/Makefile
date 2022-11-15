include .env

psqlcontainer:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -d -p ${POSTGRES_PORT}:5432 --env-file .env postgres:15-alpine

.PHONY: psqlcontainer
