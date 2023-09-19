package connections

import (
	"fmt"
	"log"
	"my_grpc/pkg/adder"
	"my_grpc/pkg/chat"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":8080"
)

type Server struct {
	Srv      *grpc.Server
	Listener net.Listener
}

func InitServerGrpc() *Server {
	server := grpc.NewServer()

	// Разрешаем клиентам получить список методов
	// reflection.Register(server)

	// Зарегистрируйте реализацию сервиса
	chat.RegisterChatServiceServer(server, &adder.ChatServer{})

	// Установите слушатель на порту 8080
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return &Server{Srv: server, Listener: listener}
}

func (s *Server) Run() {
	fmt.Println("gRPC server started...")

	// Запустите gRPC сервер
	if err := s.Srv.Serve(s.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
