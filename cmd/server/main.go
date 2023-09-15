package main

import (
	"log"
	"my_grpc/pkg/adder"
	"my_grpc/pkg/api"

	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":8080"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("cant run listener, err:%s", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
