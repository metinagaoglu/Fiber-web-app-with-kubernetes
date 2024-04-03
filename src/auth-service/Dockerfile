FROM golang:latest AS builder
WORKDIR /go/src/app
ADD . /go/src/app

RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./

RUN go mod download
RUN go build cmd/main.go

CMD ["air", "-c", ".air.toml"]