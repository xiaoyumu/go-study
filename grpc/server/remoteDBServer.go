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

func (s *server) Execute(ctx context.Context, req *rda.DbRequest) (*rda.DbResponse, error) {
	log.Println("RPC call received.")
	dumpRemoteDbRequest(req)

	return &rda.DbResponse{
		Result: "Succeeded",
	}, nil
}

func dumpRemoteDbRequest(req *rda.DbRequest) {
	log.Printf("Dumping request : Server=%v:%v;UID=%v;PWD=%v ...",
		req.Server,
		req.Port,
		req.UserId,
		req.Password)
	log.Printf("SQL:%s", req.SqlStatement)
}

func main() {
	log.Printf("Starting RPC server @ Port %v ...", port)
	tcpPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSvr := grpc.NewServer()

	// The server struct implemented the RemoteDBServiceServer interface
	rda.RegisterRemoteDBServiceServer(grpcSvr, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(grpcSvr)
	if err := grpcSvr.Serve(tcpPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
		os.Exit(-1)
	}
}
