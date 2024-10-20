# Dockerfile for Auth Service
FROM golang:1.23.2-alpine3.20  AS builder

# Set the working directory
WORKDIR /app

COPY ./grpc ./grpc
COPY ./utils ./utils
COPY ./client ./client

# Download dependencies (including those with replace directives)
RUN cd ./client && go mod download

# Build the binary
RUN cd ./client && CGO_ENABLED=0 GOOS=linux go build -o client ./cmd/grpc/main.go

# Start a new stage for a minimal image
FROM alpine:3.14

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/client/client .

# Set environment variables
# Expose the port
ENV PORT=80
EXPOSE 80

# Run the binary
CMD ["./client"]
