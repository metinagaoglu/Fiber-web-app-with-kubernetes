FROM golang:latest AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o websocket-server ./cmd/main.go
EXPOSE 8000
CMD ["./websocket-server"]

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates tzdata

# COPY --from=builder /go/src/app/ /usr/local/bin/

# WORKDIR /usr/local/bin

# CMD ["websocket-server"]
