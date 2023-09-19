package adder

import (
	"fmt"
	"my_grpc/internal/utils"
	"my_grpc/pkg/chat"
	"time"
)

const (
	SERVER_HOSTNAME string = "SERVER_HOSTNAME"
)

type ChatServer struct{}

func (s *ChatServer) SendMessage(stream chat.ChatService_SendMessageServer) error {
	isAlive := "are you alive?\n"

	dur := utils.ParserDur("10s")

	for {
		// Прочитайте сообщение от клиента
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		if req.Sender == "from_exec" {

			for i := 0; i < 5; i++ {
				fmt.Println("from_exec", req.Message)

				// Создайте ответное сообщение клиенту
				resp := &chat.MessageResponse{
					Sender:  req.Sender,
					Message: "Execute",
				}
				// Отправьте ответное сообщение клиенту
				if err := stream.Send(resp); err != nil {
					return err
				}
				time.Sleep(dur)
			}

		}

		switch req.Message {

		case "Hello Server":
			resp := &chat.MessageResponse{
				Sender:  SERVER_HOSTNAME,
				Message: isAlive,
			}
			// Отправьте ответное сообщение клиенту
			if err := stream.Send(resp); err != nil {
				return err
			}
		case "i`am ready":
			fmt.Printf("Сообщение от клиента:  %s status is ready\n", req.Sender)
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

		// Создайте ответное сообщение клиенту
		// resp := &chat.MessageResponse{
		// 	Sender:  req.Sender,
		// 	Message: "Execute",
		// }

		// Отправьте ответное сообщение клиенту
		// if err := stream.Send(resp); err != nil {
		// 	return err
		// }
	}
}
