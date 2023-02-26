# syntax=docker/dockerfile:1

## Build
FROM golang:1.20-alpine3.17 AS build

WORKDIR /app

COPY main.go go.mod ./

RUN go build -o /hello-world-api

RUN adduser -u 10001 -D nonroot

## Deploy
FROM scratch

COPY --from=build /hello-world-api /hello-world-api

COPY --from=0 /etc/passwd /etc/passwd

USER nonroot

EXPOSE 8080

ENTRYPOINT ["/hello-world-api"]