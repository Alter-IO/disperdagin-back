# Use official Golang image for building the binary
FROM golang:1.23.3-alpine AS builder

# Install timezone data
RUN apk add --no-cache tzdata

# Set timezone
ENV TZ=Asia/Makassar

# Set the working directory
WORKDIR /backend-inspektorat

# Copy go.mod and go.sum to download dependencies
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy the entire server directory
COPY server ./

# Build the Go application
RUN go build -o backend-inspektorat ./cmd/http/main.go

# Use a minimal base image for production
FROM alpine:latest

# Install timezone data and necessary libraries
RUN apk add --no-cache tzdata ca-certificates

# Set timezone
ENV TZ=Asia/Makassar

# Set the working directory
WORKDIR /app

# Copy built binary from builder stage
COPY --from=builder /backend-inspektorat/backend-inspektorat .

# Expose the application port
EXPOSE 4000

# Run the application
CMD ["./backend-inspektorat"]
