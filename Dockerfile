# syntax=docker/dockerfile:1

## Build
FROM golang:buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /ae86

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /ae86 /ae86

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/ae86"]