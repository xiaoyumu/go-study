package main

import (
	"log"
	"time"

	rda "github.com/xiaoyumu/go-study/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rda.NewRemoteDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Execute(ctx, &rda.DbRequest{
		Server:   "DBServer",
		Port:     1433,
		UserId:   "dev",
		Password: "dev2dev",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Remote Result is : %s", response.Result)
	log.Printf("Remote Message is : %s", response.Message)
}
