# syntax=docker/dockerfile:1

# Build
FROM golang:1.21 AS build
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ ./cmd/grpc-client

# Deploy
FROM alpine:3.16 as deploy
RUN mkdir /app
WORKDIR /app
COPY --from=build /app/bin/api /app/grpc-client
ENTRYPOINT [ "/app/grpc-client" ]