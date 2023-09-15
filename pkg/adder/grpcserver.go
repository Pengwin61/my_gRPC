package adder

import (
	"context"
	"my_grpc/pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) Sub(ctx context.Context, reqS *api.SubRequest) (*api.SubResponse, error) {
	return &api.SubResponse{Result: reqS.GetA() - reqS.GetB()}, nil
}

func (s *GRPCServer) Multi(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() * req.GetY()}, nil
}
