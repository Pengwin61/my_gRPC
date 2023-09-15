package adder

import (
	"context"
	"my_grpc/pkg/apiT"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *apiT.AddRequest) (*apiT.AddResponse, error) {
	return &apiT.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) Sub(ctx context.Context, reqS *apiT.SubRequest) (*apiT.SubResponse, error) {
	return &apiT.SubResponse{Result: reqS.GetA() - reqS.GetB()}, nil
}

func (s *GRPCServer) Multi(ctx context.Context, req *apiT.AddRequest) (*apiT.AddResponse, error) {
	return &apiT.AddResponse{Result: req.GetX() * req.GetY()}, nil
}
