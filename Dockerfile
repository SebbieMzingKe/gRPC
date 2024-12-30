# Use a minimal Go image
FROM golang:1.20 as builder

# Set working directory
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o kitchen ./services/kitchen/main.go
RUN go build -o orders ./services/orders/main.go

# Start a fresh image
FROM debian:bullseye-slim

# Set working directory in the container
WORKDIR /root/

# Copy the pre-built binary from the builder
COPY --from=builder /app/kitchen .
COPY --from=builder /app/orders .

# Expose ports used by the services
EXPOSE 2000 8001 9000

# Command to run multiple services
CMD ["sh", "-c", "./kitchen & ./orders"]
