package adder

import (
	"fmt"
	"my_grpc/pkg/chat"
)

type ChatServer struct{}

func (s *ChatServer) SendMessage(stream chat.ChatService_SendMessageServer) error {
	res := "are you alive?\n"

	for {
		// Прочитайте сообщение от клиента
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(req.Message)

		switch req.Message {
		case "I`m fine":
			resp := &chat.MessageResponse{
				Sender:  "1",
				Message: res,
			}
			// Отправьте ответное сообщение клиенту
			if err := stream.Send(resp); err != nil {
				return err
			}
		case "true":
			fmt.Println("case: ", req.Message)

			// case "":
			// 	fmt.Println("case: Empty")
			// default:
			// 	fmt.Println("case default")
		}

		// Создайте ответное сообщение
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
