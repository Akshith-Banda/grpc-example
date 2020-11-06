FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /example
RUN mkdir -p /example/proto
RUN mkdir -p /example/models
RUN mkdir -p /example/certs

WORKDIR /example

COPY ./proto/service.pb.go /example/proto
COPY ./main.go /example
COPY ./third_party /example
COPY ./models/setup.go /example/models
COPY ./models/tables.go /example/models
COPY ./certs/ca-cert.crt /example/certs
COPY ./certs/server-cert.crt /example/certs
COPY ./certs/ca-key.key /example/certs
COPY ./certs/server-unencrypted-key.key /example/certs

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o example .

CMD ./example