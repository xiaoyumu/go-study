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
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := rda.NewRemoteDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// dev:d3v@192.168.1.154:1433?database=godemo&connection+timeout=30
	response, err := c.ExecuteDataSet(ctx, &rda.DbRequest{
		ServerInfo: &rda.ServerInfo{
			Server:   "192.168.1.154",
			Port:     1433,
			UserId:   "dev",
			Password: "d3v",
			Database: "godemo",
		},
		SqlStatement: "SELECT GETDATE()",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Remote Result is : %v", response.Succeeded)
	log.Printf("Remote Message is : %s", response.Message)
}
