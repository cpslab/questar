FROM golang:1.17.7

WORKDIR /go/src

COPY /server/src/go.mod /go/src
COPY /server/src/go.sum /go/src

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0