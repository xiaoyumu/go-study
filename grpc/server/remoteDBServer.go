package main

import (
	"log"
	"net"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	rda "github.com/xiaoyumu/go-study/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

// NewRdaServer creates the server instance
func NewRdaServer() RdaServer {
	return &BasicRdaServer{}
}

// The RdaServer defines the basic operations
type RdaServer interface {
	ExecuteNoneQuery(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error)
	ExecuteScalar(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error)
	ExecuteDataSet(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error)
}

// BasicRdaServer implements the basic DB operation
type BasicRdaServer struct {
}

// ExecuteNoneQuery usually used for data insert/update/delete operations
// returns the effected row count.
func (s *BasicRdaServer) ExecuteNoneQuery(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error) {
	log.Printf("RPC call [ExecuteNoneQuery] received From client [%s].", s.getClietIP(ctx))
	return &rda.DBResponse{
		Succeeded: false,
		Message:   "ExecuteNoneQuery was not implemented yet",
	}, nil
}

// ExecuteScalar function perform DB operation and return the value of the first column in the first row
func (s *BasicRdaServer) ExecuteScalar(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error) {
	start := time.Now()
	log.Printf("RPC call [ExecuteScalar] received From client [%s].", s.getClietIP(ctx))

	elapsed := time.Since(start)
	log.Printf("Processed in %s", elapsed)
	return &rda.DBResponse{
		Succeeded: false,
		Message:   "ExecuteScalar was not implemented yet",
	}, nil
}

// ExecuteDataSet function usually required for complex query operation, it can be used to retrieve multiple
// result sets
func (s *BasicRdaServer) ExecuteDataSet(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error) {
	start := time.Now()
	log.Printf(">> RPC call [ExecuteDataSet] received From client [%s].", s.getClietIP(ctx))
	response := &rda.DBResponse{
		Succeeded: false,
	}

	exec := NewRdaExecutor()

	if ds, err := exec.ExecuteDataSet(req); err != nil {
		response.Message = err.Error()
	} else {
		response.Succeeded = true
		response.Dataset = ds
	}

	elapsed := time.Since(start)
	log.Printf("<< Process completed in %s", elapsed)
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

/*
func executeScalar(req *rda.DBRequest) (*rda.DBResponse, error) {
	connectionString, _ := buildConnectionString(req)
	conMgr := GetConnectionManager()
	db, err := conMgr.GetConnection(connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(req.SqlStatement)
	logAndFailOnError("executeScalar", "Prepare SQL statement failed", err)
	defer stmt.Close()

	query, err := stmt.Query()
	logAndFailOnError("executeScalar", "Query failed", err)

	columns, err := query.Columns()
	logAndFailOnError("executeScalar", "Failed to get Columns from query", err)

	log.Printf("%d columns in the query.", len(columns))

	response := rda.DBResponse{
		Succeeded: true,
		Message:   " ",
	}

	columnCount := len(columns)
	for query.Next() {
		log.Println("Processing query result...")
		// Since the query.Scan(dest ...interface{}) takes
		// a slice of pointer, we need to create two slice
		// one for actual values, and one for the pointer to
		// each actual values. Just pass the pointer slice
		// to scan method to make things work.
		values := make([]interface{}, columnCount)
		valuePtrs := make([]interface{}, columnCount)
		// Store the address of each value in values slice into
		// corresponding element of valuePtrs slice
		for i := 0; i < columnCount; i++ {
			valuePtrs[i] = &values[i]
		}

		err = query.Scan(valuePtrs...)

		if err != nil {
			log.Fatal("Scan failed:", err)
		}

		log.Println(values)

		response.ScalarValue, _ = ToDBValue(0, &values[0])
	}

	return &response, nil
}*/
