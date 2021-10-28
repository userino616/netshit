## Auth Service

This service is a part of netflix microservices which implements work with users and proxy server for grpc calls.
It provides REST API for whole service.
 
## ENV variables
Before run project you need setup next list of env variables.
You can use `.env` it will autoload on project startup. If you are going to use docker you can edit `docker-compose.yml`
```
DB_HOST=""
DB_NAME=""
DB_PORT=""
DB_USER=""
DB_PASSWORD=""

REDIS_ADDR: ""
REDIS_PASSWORD: ""

JWT_SECRET="" - secret for singing jwt tokens
JWT_ACCESS_TOKEN_EXPIRY_HOURS=""

SERVER_ADDR=""
GRPC_ADDR=""
GRPC_TIMEOUT=""

PASSWORD_SECRET="" - secret for hashing passwords

LOG_LEVEL=
```
Log level value description:
```
PanicLevel 0
FatalLevel 1
ErrorLevel 2
WarnLevel  3
InfoLevel  4
DebugLevel 5
TraceLevel 6
```

## How to run

#### Docker
1. Create docker network called `netflix` if you haven't done yet
   ```docker network create netflix```
2. Do `make`
3. Do `make migrate-up`

#### Without docker
2. Do `make migrate-up`
3. Do `make run`
###### If you are not going to use docker then you need postgresql database to be ready before launch.

### Run tests
`make test`

### Make commands
`make lint` - run linters

`make clear` - delete bin folder

`make build` - create bin file for your os

`make migrate-up / make migrate-down` - update / roll back database

`make dc-buil` - build docker image

`make dc-up` - run docker containers 

`make dc-stop` - stop docker containers

`make run` - run project without docker

`make fmt` - format go files with `gofumpt` utility

`make test` - create test db and run test then drop test db. if test fails u need do `make drop-db-test`