# syntax=docker/dockerfile:1
FROM golang:1.18.6-alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app

WORKDIR /app

ENV API_PORT=8080

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./server ./server

RUN go build -o ./app ./server/cmd 

RUN rm -rf ./server

EXPOSE $API_PORT

ENTRYPOINT ["./app"]