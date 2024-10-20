// Package server implements the gRPC server for the authentication service.
package server

import (
	"context"

	"github.com/gabehamasaki/orders/client/internal/db"
	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/zap"
)

// Server represents the gRPC server for authentication.
type Server struct {
	pb.UnimplementedClientServiceServer
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

func (s *Server) CreateClient(ctx context.Context, req *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	clientID, err := s.DB.InsertClient(ctx, db.InsertClientParams{
		Name:      req.GetName(),
		BrandName: req.GetBrandName(),
		LogoUrl:   pgtype.Text{String: req.GetLogoUrl(), Valid: req.GetLogoUrl() != ""},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateClientResponse{Id: clientID.String()}, nil
}

func (s *Server) GetClient(ctx context.Context, req *pb.GetClientRequest) (*pb.GetClientResponse, error) {
	clientID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	client, err := s.DB.FindClientById(ctx, clientID)
	if err != nil {
		return nil, err
	}

	return &pb.GetClientResponse{
		Id:        client.ID.String(),
		Name:      client.Name,
		BrandName: client.BrandName,
		LogoUrl:   client.LogoUrl.String,
	}, nil
}

func (s *Server) UpdateClient(ctx context.Context, req *pb.UpdateClientRequest) (*pb.UpdateClientResponse, error) {
	clientID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	err = s.DB.UpdateClient(ctx, db.UpdateClientParams{
		ID:        clientID,
		Name:      req.GetName(),
		BrandName: req.GetBrandName(),
		LogoUrl:   pgtype.Text{String: req.GetLogoUrl(), Valid: req.GetLogoUrl() != ""},
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateClientResponse{}, nil
}

func (s *Server) DeleteClient(ctx context.Context, req *pb.DeleteClientRequest) (*pb.DeleteClientResponse, error) {
	clientID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}

	err = s.DB.DeleteClient(ctx, clientID)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteClientResponse{}, nil
}

func (s *Server) ListClients(ctx context.Context, req *pb.ListClientsRequest) (*pb.ListClientsResponse, error) {
	clients, err := s.DB.GetClients(ctx, db.GetClientsParams{
		Limit:  req.GetPerPage(),
		Offset: req.GetPage() - 1,
	})
	if err != nil {
		return nil, err
	}

	var pbClients []*pb.ListClientsResponse_Clients
	var total int32
	var totalPage int32
	for _, client := range clients {
		total = int32(client.Total)
		totalPage = int32(client.TotalPages)
		pbClients = append(pbClients, &pb.ListClientsResponse_Clients{
			Id:        client.ID.String(),
			Name:      client.Name,
			BrandName: client.BrandName,
			LogoUrl:   client.LogoUrl.String,
		})
	}

	return &pb.ListClientsResponse{Clients: pbClients, Total: total, TotalPages: totalPage}, nil
}
