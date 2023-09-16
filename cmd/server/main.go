package main

import connections "my_grpc/internal/connection"

func main() {

	s := connections.InitServerGrpc()

	s.Run()

}
