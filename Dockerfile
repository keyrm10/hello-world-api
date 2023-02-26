# syntax=docker/dockerfile:1

## Build
FROM golang:1.20-alpine3.17 AS build

WORKDIR /app

COPY main.go go.mod ./

RUN go build -o /hello-world-api

## Deploy
FROM scratch

COPY --from=build /hello-world-api /hello-world-api

EXPOSE 8080

ENTRYPOINT ["/hello-world-api"]