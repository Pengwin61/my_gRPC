package main

import (
	"fmt"
	"log"
	connections "my_grpc/internal/connection"
)

const (
	CLIENT_HOSTNAME string = "CLIENT_ACTOR-01"
	strRes          string = "96015|anton.sadym-50-1694950510_stDMATE_dp32|50|mk0vm1067.bosch-ru.ru|R|2023-09-17T14:35:10|5f65727b9a8646548578f290c22e4d9a|172.22.43.46|38773|38774|2023-09-17T14:35:11|anton.sadym|18966|38775|-1|-1"
)

func main() {

	c := connections.InitClientGrpc()

	defer c.Conn.Close()

	// dur := utils.ParserDur("10s")

	initMessage := true

	for {

		// Первое сообщение серверу Init Message
		if initMessage {
			c.SendMsg(CLIENT_HOSTNAME, "Hello Server")
			initMessage = false
		}

		msg := c.ReadMsg()

		fmt.Println("msg:", msg)
		switch msg {

		case "Execute":
			// result := utils.Execut()
			// c.SendMsg("from_exec", result)
			c.SendMsg("from_exec", strRes)

		case "are you alive?\n":
			c.SendMsg("macbook-air", "i`am ready")

		case "terminated":
			fmt.Println("Session terminated")

		}

		log.Println("Конец консольного цикла")

		// time.Sleep(dur)
	}

}
