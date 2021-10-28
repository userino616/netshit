APP_NAME := "netflix_movies"
DB_CONTAINER_NAME := "netflix_movies_postgres"
DB_NAME := "netflix_movies"
DB_PORT := "5477"

all: dc-build dc-up

lint:
	golangci-lint run --fix -c linters.yml

clear:
	rm -r -f ./bin

build:
	go build -o ./bin/main ./cmd/main

migrate-up:
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

dc-build:
	docker-compose build ${APP_NAME}

dc-up:
	docker-compose up -d

dc-stop:
	docker-compose stop

run:
	go run ./cmd/main

fmt:
	gofumpt -l -w .

create-db-test:
	docker exec -it ${DB_CONTAINER_NAME} createdb --username=postgres --owner=postgres netflix_test

up-db-test: create-db-test
	migrate -path sql/migrations -database "postgresql://postgres:qwerty@localhost:5477/netflix_test?sslmode=disable" -verbose up

drop-db-test:
	docker exec -it ${DB_CONTAINER_NAME} dropdb --username=postgres netflix_test

test: up-db-test
	go test -race -v ./... && make drop-db-test
