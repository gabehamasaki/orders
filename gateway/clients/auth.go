package clients

import (
	"context"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	conn, err := grpc.NewClient(c.cfg.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.logger.Error("failed to dial auth service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)

	return client.Register(ctx, request)
}

func (c *Client) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	conn, err := grpc.NewClient(c.cfg.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.logger.Error("failed to dial auth service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)

	return client.Login(ctx, request)
}

func (c *Client) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	conn, err := grpc.NewClient(c.cfg.AuthServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.logger.Error("failed to dial auth service", zap.Error(err))
		return nil, err
	}
	defer conn.Close()
	client := pb.NewAuthServiceClient(conn)

	return client.ValidateToken(ctx, request)
}
