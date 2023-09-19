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
	Conn       *grpc.ClientConn
	MainStream chat.ChatService_SendMessageClient
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
	// streamH, _ := client.Hello(context.Background())

	return &Conn{Conn: conn, MainStream: stream}
}

func (c *Conn) SendMsg(to, message string) {

	req := &chat.MessageRequest{
		Sender:  to,
		Message: message,
	}

	if err := c.MainStream.Send(req); err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

}

func (c *Conn) ReadMsg() string {
	// Прочитайте ответное сообщение от сервера
	resp, err := c.MainStream.Recv()
	if err != nil {
		log.Fatalf("failed to receive message: %v", err)
	}

	// fmt.Printf("Received response: %s\n", resp.Message)
	fmt.Printf("Сообщение от сервера: %s\n", resp.Message)

	return resp.Message
}
