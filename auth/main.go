package main

import (
	"context"
	"fmt"
	"net"

	"github.com/bufbuild/protovalidate-go"
	"github.com/gabehamasaki/orders/auth/db"
	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logger *zap.Logger

type Server struct {
	pb.UnimplementedAuthServiceServer
	db     *db.Queries
	logger *zap.Logger
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	if err := v.Validate(req); err != nil {
		return nil, err
	}

	passBytes, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), 15)
	if err != nil {
		return nil, err
	}

	id, err := s.db.InserUser(ctx, db.InserUserParams{
		Email:    pgtype.Text{String: req.GetEmail(), Valid: true},
		Password: pgtype.Text{String: string(passBytes), Valid: true},
		Name:     pgtype.Text{String: req.GetName(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Token: "token",
		Id:    id.String(),
	}, nil
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

	listener, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		panic(err)
	}

	logger.Info("Server starting in port :10001")
	server := grpc.NewServer(grpc.UnaryInterceptor(logRequest))
	pb.RegisterAuthServiceServer(server, &Server{
		db:     db.New(pool),
		logger: logger,
	})
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

func logRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger.Info(fmt.Sprintf("method %q", info.FullMethod))
	resp, err := handler(ctx, req)
	if err != nil {
		logger.Error(fmt.Sprintf("method %q failed: %s", info.FullMethod, err.Error()))
		return resp, status.Errorf(codes.Internal, "%v", err)
	}

	return resp, err
}
