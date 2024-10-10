package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

var logger *zap.Logger

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return nil, fmt.Errorf("Testing...")
}

func main() {
	ctx := context.Background()
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = zapLogger.Named("auth-service")
	defer logger.Sync()

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s", "postgres", "password", "localhost", "5432", "orders"))
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	logger.Info("Server starting in port :9001")
	server := grpc.NewServer(grpc.UnaryInterceptor(logRequest))
	pb.RegisterAuthServiceServer(server, &Server{})
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

func logRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info(fmt.Sprintf("method %q", info.FullMethod))
	resp, err := handler(ctx, req)
	if err != nil {
		logger.Error(fmt.Sprintf("method %q failed: %s", info.FullMethod, err.Error()))
	}
	return resp, err
}
