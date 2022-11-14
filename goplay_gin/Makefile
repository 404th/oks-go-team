#!make
include .env

# postgres:
	# docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -p ${POSTGRES_PORT}:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:15-alpine

postgres:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -d -p ${POSTGRES_PORT}:5432 --env-file .env postgres:15-alpine

createdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} dropdb ${POSTGRES_DB} --username=${POSTGRES_USER}

opendb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

migrateup:
	migrate -path ./db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${PGSSLMODE}" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${PGSSLMODE}" -verbose down

swag:
	swag init -g cmd/main.go

test:
	go test -v -cover ./...

.PHONY: postgres, createdb, dropdb, opendb, migrateup, migratedown, swag, test
