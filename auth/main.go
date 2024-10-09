package main

import (
	"context"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func main() {
}
