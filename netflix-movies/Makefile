APP_NAME := "netflix-movies"
DB_CONTAINER_NAME := "netflix_postgres"
DB_NAME := "netflix"

all: migrate-up dc-build dc-up

lint:
	golangci-lint run --fix -c linters.yml

clear:
	rm -r -f ./bin

build:
	go build -o ./bin/main ./cmd/main

create-migrations:
	migrate create -ext sql -dir sql/migrations -seq init_schema

migrate-up:
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

dc-build:
	docker build -t movies-service .

dc-up:
	docker run -d --rm  --name movies-service --network=netflix-auth_default movies-service

run:
	go run ./cmd/main

fmt:
	gofumpt -l -w .

gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./src/grpc/proto/movie.proto

create-db-test:
	docker exec -it netflix_postgres createdb --username=postgres --owner=postgres netflix_test

up-db-test: create-db-test
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5432/netflix_test?sslmode=disable" -verbose up

drop-db-test:
	docker exec -it netflix_postgres dropdb --username=postgres netflix_test

test: up-db-test
	go test -race -v ./... && make drop-db-test
