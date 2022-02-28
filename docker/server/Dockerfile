FROM golang:1.17.7

WORKDIR /go/src

COPY go.mod /go/src
COPY go.sum /go/src

RUN go mod download