APP_NAME := "netflix-auth"
DB_CONTAINER_NAME := "netflix_postgres"
DB_NAME := "netflix"

all: dc-build dc-up

clear:
	rm -r -f ./bin

lint:
	golangci-lint run --fix -c linters.yml

build:
	go build -o ./bin/main ./cmd/main

dc-build:
	docker-compose build proxy_server

dc-up:
	docker-compose up -d

dc-stop:
	docker-compose stop

create-migrations:
	migrate create -ext sql -dir sql/migrations -seq init_schema

create-db:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=postgres --owner=postgres ${DB_NAME}

drop-db:
	docker exec -it ${DB_CONTAINER_NAME} dropdb --username=postgres ${DB_NAME}

migrate-up:
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

fmt:
	gofumpt -l -w .

create-db-test:
	docker exec -it netflix_postgres createdb --username=postgres --owner=postgres netflix_test

up-db-test: create-db-test
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5432/netflix_test?sslmode=disable" -verbose up

drop-db-test:
	docker exec -it netflix_postgres dropdb --username=postgres netflix_test

test: up-db-test
	go test -race -v ./internal/... && make drop-db-test