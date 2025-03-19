# Builder
FROM golang:1.24.0 AS builder
LABEL authors="virgiawanly"

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o build/server "./cmd/server/"

# Test Runner
FROM builder AS test-runner

RUN go test ./...

# Runner
FROM alpine:latest AS runner

# Copy binary
COPY --from=builder /app/build/* /app/

# Copy migrations
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/makefile /app/

WORKDIR /app

RUN apk add --no-cache make curl tzdata
RUN curl -sSL https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.13 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

# Grpc Server
EXPOSE 8080

ENTRYPOINT ["./server"]