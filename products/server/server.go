// Package server implements the gRPC server for the authentication service.
package server

import (
	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gabehamasaki/orders/products/db"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/zap"
)

// Server represents the gRPC server for authentication.
type Server struct {
	pb.UnimplementedProductServiceServer
	DB     *db.Queries
	Logger *zap.Logger
}

// NewServer creates a new instance of the authentication server.
func NewServer(logger *zap.Logger, pool *pgxpool.Pool) *Server {
	return &Server{
		DB:     db.New(pool),
		Logger: logger,
	}
}
