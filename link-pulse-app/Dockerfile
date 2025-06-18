# Stage 1: Builder with Go 1.24.3
FROM golang:1.24.3-alpine AS builder

# Set environment
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create app directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project source
COPY . .

# Build the Go binary
RUN go build -o main main.go

# Stage 2: Minimal runtime container
FROM alpine:latest

# Create non-root user for security
RUN adduser -D appuser

WORKDIR /app

# Copy built binary from builder stage
COPY --from=builder /app/main .

# Use non-root user
USER appuser

# Expose the port your app listens on (change if needed)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./main"]
