# Use Go official image
FROM golang:1.23.3-alpine

# Install timezone data
RUN apk add --no-cache tzdata

# Set timezone into Asia/Makassar
ENV TZ=Asia/Makassar

# Set the working directory
WORKDIR /backend-alter-io

# Copy go.mod and go.sum to download dependencies
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy the entire server directory
COPY server ./

# Install Air for live reload (optional, remove if not needed)
RUN go install github.com/air-verse/air@latest

# Expose the app port
EXPOSE 4000

# Set the default command
CMD ["air"]
