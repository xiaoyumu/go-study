package main

import (
	"log"
	"net"

	rda "github.com/xiaoyumu/go-study/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	log.Printf("Starting RPC server @ Port %v ...", port)
	tcpPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSvr := grpc.NewServer()

	// The server struct implemented the RemoteDBServiceServer interface
	rda.RegisterRemoteDBServiceServer(grpcSvr, NewRdaServer())
	// Register reflection service on gRPC server.
	reflection.Register(grpcSvr)
	if err := grpcSvr.Serve(tcpPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
