// Package server implements the gRPC server for the authentication service.
package server

import (
	"context"
	"fmt"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gabehamasaki/orders/products/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	_, err := s.DB.InsertProduct(ctx, db.InsertProductParams{
		Name:        req.GetName(),
		Description: pgtype.Text{String: req.GetDescription(), Valid: true},
		Price:       req.GetPrice(),
		ImageUrl:    pgtype.Text{String: req.GetImageUrl(), Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{}, nil
}
func (s *Server) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := s.DB.GetProducts(ctx, db.GetProductsParams{
		Limit:  req.GetPerPage(),
		Offset: req.GetPage() - 1,
	})
	if err != nil {
		return nil, err
	}

	var pbProducts []*pb.ListProductsResponse_Product
	var total int32
	var totalPage int32
	for _, product := range products {
		total = int32(product.Total)
		totalPage = int32(product.TotalPages)
		pbProducts = append(pbProducts, &pb.ListProductsResponse_Product{
			Id:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description.String,
			Price:       product.Price,
			ImageUrl:    product.ImageUrl.String,
		})
	}

	return &pb.ListProductsResponse{Products: pbProducts, Total: total, TotalPages: totalPage}, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	ID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	product, err := s.DB.GetProduct(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &pb.GetProductResponse{
		Id:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description.String,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl.String,
	}, nil

}
