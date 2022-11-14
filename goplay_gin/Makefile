postgres:
	docker run --name bookpostgres -p 2020:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=amazingsecretkeyfordatabase -d postgres:15-alpine

createdb:
	docker exec -it bookpostgres createdb --username=postgres --owner=postgres bookpostgres

dropdb:
	docker exec -it bookpostgres dropdb bookpostgres --username=postgres

opendb:
	docker exec -it bookpostgres psql -U postgres -d bookpostgres

migrateup:
	migrate -path ./db/migrations -database "postgresql://postgres:amazingsecretkeyfordatabase@localhost:2020/bookpostgres?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migrations -database "postgresql://postgres:amazingsecretkeyfordatabase@localhost:2020/bookpostgres?sslmode=disable" -verbose down

swag:
	swag init -g cmd/main.go

test:
	go test -v -cover ./...

.PHONY: postgres, createdb, dropdb, opendb
