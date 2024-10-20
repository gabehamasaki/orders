# Dockerfile for Gateway Service
FROM golang:1.23.2-alpine3.20  AS builder

# Set the working directory
WORKDIR /app

COPY ./grpc ./grpc
COPY ./utils ./utils
COPY ./gateway ./gateway

# Download dependencies (including those with replace directives)
RUN cd ./gateway && go mod download

# Build the binary
RUN cd ./gateway && CGO_ENABLED=0 GOOS=linux go build -o gateway ./cmd/web/main.go

# Start a new stage for a minimal image
FROM alpine:3.14

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gateway/gateway .

# Set environment variables
ENV PORT=80
EXPOSE 80

# Run the binary
CMD ["./gateway"]
