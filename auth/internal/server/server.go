// Package server implements the gRPC server for the authentication service.
package server

import (
	"context"

	"github.com/gabehamasaki/orders/auth/internal/db"
	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represents the gRPC server for authentication.
type Server struct {
	pb.UnimplementedAuthServiceServer
	DB        *db.Queries
	Logger    *zap.Logger
	SecretKey []byte
}

// NewServer creates a new instance of the authentication server.
func NewServer(logger *zap.Logger, pool *pgxpool.Pool, SecretKey []byte) *Server {
	return &Server{
		DB:        db.New(pool),
		Logger:    logger,
		SecretKey: SecretKey,
	}
}

// Login handles user authentication and returns a JWT token upon successful login.
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Find user by email
	user, err := s.DB.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// Compare provided password with stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid credentials")
	}
	// Create JWT token
	token, err := s.createToken(user.ID.String(), uuid.Must(uuid.FromBytes(user.ClientID.Bytes[:])).String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create token: %v", err)
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil
}

// ValidateToken verifies the provided JWT token and returns the associated user ID if valid.
func (s *Server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	payload, err := s.verifyToken(ctx, req.GetToken())
	if err != nil {
		return &pb.ValidateTokenResponse{
			Id:    "",
			Valid: false,
		}, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
	}
	return &pb.ValidateTokenResponse{
		Id:       payload.UserID,
		ClientId: payload.ClientID,
		Valid:    true,
	}, nil
}
