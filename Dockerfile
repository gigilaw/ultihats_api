# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /api

COPY . .
RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 3000

ENTRYPOINT CompileDaemon -command="./ultihats"