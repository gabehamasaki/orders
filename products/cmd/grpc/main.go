// Package main is the entry point for the authentication service.
package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gabehamasaki/orders/products/internal/config"
	"github.com/gabehamasaki/orders/products/internal/server"
	log "github.com/gabehamasaki/orders/utils/logger"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

// Global logger instance
var (
	logger *zap.Logger
)

// main is the entry point of the application.
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// run initializes and starts the gRPC server.
func run() error {
	ctx := context.Background()

	// Initialize logger
	var err error

	// Configure and build the logger
	logger, err = log.NewLogger("product")
	if err != nil {
		return err
	}
	defer logger.Sync()

	// Load configuration from environment variables
	config, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	// Initialize database connection pool
	pool, err := pgxpool.New(ctx, config.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer pool.Close()

	// Verify database connection
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Initialize gRPC server
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	logger.Info("Starting TCP server", zap.String("address", config.ServerAddress))
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(logAndVerifyAnyRequest))

	// Register Auth Service
	pb.RegisterProductServiceServer(grpcServer, server.NewServer(logger, pool))
	return grpcServer.Serve(listener)
}

// logAndVerifyAnyRequest is a gRPC interceptor that logs requests and validates them.
func logAndVerifyAnyRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	logger.Info("Received request", zap.String("method", info.FullMethod))

	// Call the handler
	resp, err := handler(ctx, req)
	duration := time.Since(start)

	// Log errors if any
	if err != nil {
		logger.Error("Request failed",
			zap.String("method", info.FullMethod),
			zap.Duration("duration", duration),
			zap.Error(err),
		)
		return nil, err
	}

	// Log successful requests
	logger.Info("Request completed",
		zap.String("method", info.FullMethod),
		zap.Duration("duration", duration),
	)

	return resp, nil
}
