# Stage 1: Build
FROM golang:1.23.4 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main cmd/main.go

# Stage 2: Runtime
FROM gcr.io/distroless/base-debian12:latest

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

#Set to release mode for production
ENV GIN_MODE=release

# Start the application
CMD ["./main"]
