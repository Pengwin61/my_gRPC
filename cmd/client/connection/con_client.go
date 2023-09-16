package connections

import (
	"context"
	"fmt"
	"log"
	"my_grpc/pkg/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Conn struct {
	Conn *grpc.ClientConn
	// Client chat.ChatServiceClient
	Stream chat.ChatService_SendMessageClient
}

func InitClientGrpc() *Conn {

	// Установите соединение с сервером
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// Создайте клиентский объект
	client := chat.NewChatServiceClient(conn)

	// Создайте поток для отправки сообщений на сервер
	stream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	return &Conn{Conn: conn, Stream: stream}
}

func (c *Conn) SendMessages(to, message string) {

	req := &chat.MessageRequest{
		Sender:  to,
		Message: message,
	}

	if err := c.Stream.Send(req); err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

}

func (c *Conn) ReadMessages() string {
	// Прочитайте ответное сообщение от сервера
	resp, err := c.Stream.Recv()
	if err != nil {
		log.Fatalf("failed to receive message: %v", err)
	}

	// fmt.Printf("Received response: %s\n", resp.Message)
	fmt.Printf("Ответ от сервера: %s\n", resp.Message)

	return resp.Message
}
