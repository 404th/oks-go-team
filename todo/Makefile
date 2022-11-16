include .env

psqlcontainer:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -d -p ${POSTGRES_PORT}:5432 --env-file .env postgres:15-alpine

migrateup:
	migrate -path ./schema -database "postgresql://postgres:${POSTGRES_PASSWORD}@localhost:21212/storage?sslmode=disable" -verbose up

migratedown:
	migrate -path ./schema -database "postgresql://postgres:${POSTGRES_PASSWORD}@localhost:21212/storage?sslmode=disable" -verbose down

.PHONY: psqlcontainer, migrateup, migratedown
