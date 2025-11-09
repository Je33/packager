# Multi-stage build for Go backend
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /build

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with security flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=$(git rev-parse --short HEAD 2>/dev/null || echo 'dev')" \
    -trimpath \
    -o api \
    ./cmd/api/main.go

# Runtime stage - distroless for minimal attack surface
FROM gcr.io/distroless/static-debian12:nonroot

# Copy timezone data and CA certificates from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /build/api /app/api

# Use non-root user (distroless nonroot is UID 65532)
USER nonroot:nonroot

WORKDIR /app

# Expose port
EXPOSE 8080

# Run
ENTRYPOINT ["/app/api"]
