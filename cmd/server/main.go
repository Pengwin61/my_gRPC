package main

import (
	"log"
	"my_grpc/pkg/adder"
	"my_grpc/pkg/api"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":8080"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	// Разрешаем клиентам получить список методов
	reflection.Register(s)

	api.RegisterAdderServer(s, srv)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("cant run listener, err:%s", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
