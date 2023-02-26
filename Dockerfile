# syntax=docker/dockerfile:1

FROM golang:1.20-alpine3.17 AS build

WORKDIR /app

COPY main.go go.mod ./

RUN go build -o /hello-world-api

EXPOSE 8080

CMD ["/hello-world-api"]