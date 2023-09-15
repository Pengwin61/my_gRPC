package main

import (
	"context"
	"fmt"
	"log"
	"my_grpc/pkg/apiT"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	c := apiT.NewAdderClient(conn)

	res, err := c.Add(context.Background(), &apiT.AddRequest{X: int32(10), Y: int32(10)})
	if err != nil {
		fmt.Println(err)
	}

	log.Println("res:", res)

}
