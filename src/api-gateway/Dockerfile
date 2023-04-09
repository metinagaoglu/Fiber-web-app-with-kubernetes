FROM golang:1.20-alpine AS build-stage
WORKDIR /app
ADD . /app
RUN go mod download
RUN go build cmd/main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main /main
COPY --from=build-stage /app/pkg/config/envs/dev.env /pkg/config/envs/dev.env

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/main"]
