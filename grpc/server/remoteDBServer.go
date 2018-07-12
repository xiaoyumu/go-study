package main

import (
	"log"
	"net"
	"os"

	rda "github.com/xiaoyumu/go-study/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Execute(context.Context, *rda.DbRequest) (*rda.DbResponse, error) {
	log.Println("RPC call received.")
	return &rda.DbResponse{
		Result: "Succeeded",
	}, nil
}

func main() {
	log.Printf("Starting RPC server @ Port %v ...", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rda.RegisterRemoteDBServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		os.Exit(-1)
	}
}
