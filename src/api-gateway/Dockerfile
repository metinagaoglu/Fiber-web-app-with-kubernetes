FROM golang:latest AS builder
WORKDIR /go/src/app
ADD . /go/src/app

RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./

RUN go mod download
RUN go build cmd/main.go

CMD ["air", "-c", ".air.toml"]

# Deploy the application binary into a lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage

# WORKDIR /

# COPY --from=build-stage /app/main /main
# COPY --from=build-stage /app/pkg/config/envs/dev.env /pkg/config/envs/dev.env

# EXPOSE 3000

# USER nonroot:nonroot

#ENTRYPOINT ["/main"]
