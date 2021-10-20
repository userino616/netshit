# image for building application
FROM golang:1.17-alpine as builder
RUN apk add --no-cache make g++
WORKDIR /src
ADD . .
RUN make clear build

# image for running application
FROM alpine:latest
WORKDIR /app
COPY config.yml .
COPY --from=builder ./src/bin .
EXPOSE 8000
ENTRYPOINT ["./main"]