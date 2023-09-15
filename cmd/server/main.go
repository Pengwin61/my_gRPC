package main

import connections "my_grpc/cmd/client/connection"

func main() {

	s := connections.InitServerGrpc()

	s.Run()

}
