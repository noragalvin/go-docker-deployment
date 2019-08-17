FROM golang:1.12-alpine

ADD . /app

WORKDIR /app

RUN apk add bash ca-certificates git gcc g++ libc-dev

RUN CGO_ENABLED=0 GOOS=linux

ARG GO111MODULE=on

RUN go build -a -installsuffix .

EXPOSE 8000

ENTRYPOINT ./go-docker-development