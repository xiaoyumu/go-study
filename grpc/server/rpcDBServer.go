package main

import (
	"log"
	"net"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xiaoyumu/go-study/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

// NewRdaServer creates the server instance
func NewRdaServer() RPCDBProxyServer {
	return &BasicRdaServer{}
}

// The RPCDBProxyServer defines the basic operations
type RPCDBProxyServer interface {
	ExecuteNoneQuery(ctx context.Context, req *proto.DBRequest) (*proto.DBResponse, error)
	ExecuteScalar(ctx context.Context, req *proto.DBRequest) (*proto.DBResponse, error)
	ExecuteDataSet(ctx context.Context, req *proto.DBRequest) (*proto.DBResponse, error)
}

// BasicRdaServer implements the basic DB operation
type BasicRdaServer struct {
}

// ExecuteNoneQuery usually used for data insert/update/delete operations
// returns the effected row count.
func (s *BasicRdaServer) ExecuteNoneQuery(ctx context.Context, req *proto.DBRequest) (*proto.DBResponse, error) {
	start := time.Now()
	response := &proto.DBResponse{Succeeded: false}

	exec := NewRdaExecutor()

	if rowEffected, err := exec.ExecuteNoneQuery(req); err != nil {
		response.Message = err.Error()
	} else {
		response.Succeeded = true
		response.RowEffected = rowEffected
	}

	elapsed := time.Since(start)
	defer log.Printf("<<   RPC [ExecuteNoneQuery] completed in [ %s ] ClientIP: [%s]", elapsed, s.getClietIP(ctx))
	return response, nil
}

// ExecuteScalar function perform DB operation and return the value of the first column in the first row
func (s *BasicRdaServer) ExecuteScalar(ctx context.Context, req *proto.DBRequest) (*proto.DBResponse, error) {
	start := time.Now()
	response := &proto.DBResponse{Succeeded: false}

	exec := NewRdaExecutor()

	if sv, err := exec.ExecuteSalar(req); err != nil {
		response.Message = err.Error()
	} else {
		response.Succeeded = true
		response.ScalarValue = sv
	}

	elapsed := time.Since(start)
	defer log.Printf("<<   RPC [ExecuteScalar] completed in [ %s ] ClientIP: [%s]", elapsed, s.getClietIP(ctx))
	return response, nil
}

// ExecuteDataSet function usually required for complex query operation, it can be used to retrieve multiple
// result sets
func (s *BasicRdaServer) ExecuteDataSet(ctx context.Context, req *proto.DBRequest) (*proto.DBResponse, error) {
	start := time.Now()
	response := &proto.DBResponse{Succeeded: false}

	exec := NewRdaExecutor()

	if ds, err := exec.ExecuteDataSet(req); err != nil {
		response.Message = err.Error()
	} else {
		response.Succeeded = true
		response.Dataset = ds
	}

	elapsed := time.Since(start)
	defer log.Printf("<<   RPC [ExecuteDataSet] completed in [ %s ] ClientIP: [%s]", elapsed, s.getClietIP(ctx))
	return response, nil
}

func (s *BasicRdaServer) getClietIP(ctx context.Context) string {
	pr, ok := peer.FromContext(ctx)

	if !ok {
		log.Println("[getClinetIP] invoke FromContext() failed")
		return ""
	}
	if pr.Addr == net.Addr(nil) {
		log.Println("[getClientIP] peer.Addr is nil")
		return ""
	}

	addSlice := strings.Split(pr.Addr.String(), ":")

	return addSlice[0]
}

func logAndFailOnError(prefix string, description string, err error) {
	if err != nil {
		log.Printf("[%s] %s: %s", prefix, description, err)
		panic(err)
	}
}
