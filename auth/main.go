// Package main is the entry point for the authentication service.
package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/gabehamasaki/orders/auth/config"
	"github.com/gabehamasaki/orders/auth/server"
	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger, err := cfg.Build()
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	logger = zapLogger.Named("AuthService")
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

	logger.Info("Server starting", zap.String("address", config.ServerAddress))
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(logAndVerifyAnyRequest))

	// Register Auth Service
	pb.RegisterAuthServiceServer(grpcServer, server.NewServer(logger, pool, []byte(config.JWTSecretKey)))
	return grpcServer.Serve(listener)
}

// logAndVerifyAnyRequest is a gRPC interceptor that logs requests and validates them.
func logAndVerifyAnyRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	logger.Info("Received request", zap.String("method", info.FullMethod))

	// Initialize request validator
	v, err := protovalidate.New()
	if err != nil {
		logger.Error("Failed to initialize validator", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to initialize validator: %v", err)
	}

	// Validate the request if it's a protobuf message
	if protoReq, ok := req.(protoreflect.ProtoMessage); ok {
		if err := v.Validate(protoReq); err != nil {
			logger.Error("Invalid request",
				zap.String("method", info.FullMethod),
				zap.Error(err),
			)
			return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
		}
	}

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
