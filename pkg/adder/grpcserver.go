package adder

import (
	"fmt"
	"my_grpc/pkg/chat"
)

type ChatServer struct{}

func (s *ChatServer) SendMessage(stream chat.ChatService_SendMessageServer) error {
	isAlive := "are you alive?\n"

	for {
		// Прочитайте сообщение от клиента
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(req.Message)

		switch req.Message {

		case "Hello":
			resp := &chat.MessageResponse{
				Sender:  "server",
				Message: isAlive,
			}
			// Отправьте ответное сообщение клиенту
			if err := stream.Send(resp); err != nil {
				return err
			}

		}

		// Создайте ответное сообщение клиенту
		resp := &chat.MessageResponse{
			Sender:  req.Sender,
			Message: "Execute",
		}

		// Отправьте ответное сообщение клиенту
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func (s *ChatServer) Hello(stream chat.ChatService_HelloServer) error {
	for {
		// Прочитайте сообщение от клиента
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(req.Message, req.Name)

		// Создайте ответное сообщение клиенту
		resp := &chat.MsgHelloResponse{
			Name:    req.Name,
			Message: "Pong",
		}

		// Отправьте ответное сообщение клиенту
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}
