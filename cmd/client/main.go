package main

import (
	"context"
	"fmt"
	"log"
	"my_grpc/pkg/api"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	c := api.NewAdderClient(conn)

	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(10), Y: int32(10)})
	if err != nil {
		fmt.Println(err)
	}

	log.Println("res:", res)

}
