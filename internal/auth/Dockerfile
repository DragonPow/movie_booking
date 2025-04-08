# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install required build tools
RUN apk add --no-cache make git

# Copy only go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the auth service
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./internal/auth/cmd/server

# Final stage
FROM alpine:3.18

WORKDIR /app

# Install certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /auth-service .

# Copy migration files
COPY --from=builder /app/internal/auth/migrations ./migrations/

# Copy config files
COPY --from=builder /app/internal/auth/config.yaml ./config.yaml
COPY --from=builder /app/internal/auth/.env ./.env

EXPOSE 50051

# Command to run the service with specific config file
CMD ["./auth-service", "--config", "config.yaml"]