# image for building application
FROM golang:1.17-alpine as builder
RUN apk add --no-cache make g++
WORKDIR /src
ADD . .
RUN make clear build

# image for running application
FROM alpine:latest
ENV DB_HOST="netflix_users_postgres"
ENV DB_NAME="netflix_users"
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD="qwerty"

ENV SERVER_ADDR=":8085"
ENV LOG_LEVEL=6

WORKDIR /app
COPY --from=builder ./src/bin .
EXPOSE 8085
ENTRYPOINT ["./main"]