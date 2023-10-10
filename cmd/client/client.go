package main

import (
	"fmt"
	connections "my_grpc/internal/connection"
)

const (
	CLIENT_HOSTNAME string = "CLIENT_ACTOR-01"
	strRes          string = "anton.sadym-50-1694950510_stDMATE_dp32|mk0vm1067.bosch-ru.ru|R|2023-09-17T14:35:10|172.22.43.46|2023-09-17T14:35:11|anton.sadym|"
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

		switch msg {

		case "Execute":
			// result := utils.Execut()
			// c.SendMsg("from_exec", result)
			c.SendMsg("from_exec", strRes)

		case "are you alive?\n":
			c.SendMsg(CLIENT_HOSTNAME, "i`am ready")

		case "terminated":
			fmt.Println("Session terminated")

		}

		fmt.Println("#------------------------------------------------#")

		// time.Sleep(dur)
	}

}
