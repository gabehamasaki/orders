package clients

import (
	"context"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	conn, err := grpc.NewClient(c.cfg.ProductsServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.logger.Error("failed to dial auth service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)

	return client.CreateProduct(ctx, request)
}

func (c *Client) ListProducts(ctx context.Context, request *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	conn, err := grpc.NewClient(c.cfg.ProductsServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.logger.Error("failed to dial auth service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)

	return client.ListProducts(ctx, request)
}
