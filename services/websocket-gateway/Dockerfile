FROM golang:latest AS builder

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download


#RUN go build -o websocket-server ./cmd/main.go

#CMD ["./websocket-server"]
#CMD ["make","server"]

CMD ["air", "-c", ".air.toml"]
# FROM alpine:latest

# RUN apk --no-cache add ca-certificates tzdata

# COPY --from=builder /go/src/app/ /usr/local/bin/

# WORKDIR /usr/local/bin

# CMD ["websocket-server"]
