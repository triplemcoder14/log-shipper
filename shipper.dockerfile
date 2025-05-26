FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o log-shipper main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/log-shipper .

COPY configs ./configs

ENV OTEL_CONFIG_FILE=/app/configs/collector-config.yaml

EXPOSE 4317

ENTRYPOINT ["./log-shipper", "--config", "/app/configs/collector-config.yaml"]
