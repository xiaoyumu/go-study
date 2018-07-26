package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	rda "github.com/xiaoyumu/go-study/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) ExecuteNoneQuery(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error) {
	log.Printf("RPC call [ExecuteNoneQuery] received From client [%s].", getClietIP(ctx))
	return &rda.DBResponse{
		Succeeded: false,
		Message:   "ExecuteNoneQuery was not implemented yet",
	}, nil
}

func (s *server) ExecuteScalar(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error) {
	log.Printf("RPC call [ExecuteScalar] received From client [%s].", getClietIP(ctx))

	return &rda.DBResponse{
		Succeeded: false,
		Message:   "ExecuteScalar was not implemented yet",
	}, nil
}

func (s *server) ExecuteDataSet(ctx context.Context, req *rda.DBRequest) (*rda.DBResponse, error) {
	log.Printf("RPC call [ExecuteDataSet] received From client [%s].", getClietIP(ctx))
	response := &rda.DBResponse{
		Succeeded: false,
	}

	if ds, err := executeDataSet(req); err != nil {
		response.Message = err.Error()		
	} else {
		response.Succeeded = true
		response.Dataset = ds
	}

	return response, nil
}

func getClietIP(ctx context.Context) string {
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

func dumpRemoteDbRequest(req *rda.DBRequest) {
	connStr, _ := buildConnectionString(req)

	log.Printf("Dumping request : %s ...", connStr)
	log.Printf("SQL:%s", req.SqlStatement)
}

func buildConnectionString(req *rda.DBRequest) (string, error) {
	// Sample Connection string:
	// sqlserver://dev:d3v@192.168.1.154:1433?database=godemo&connection+timeout=30
	conn := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s&connection+timeout=30",
		req.ServerInfo.UserId,
		req.ServerInfo.Password,
		req.ServerInfo.Server,
		req.ServerInfo.Port,
		req.ServerInfo.Database)
	return conn, nil
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

// Execute the remote Db request 
func executeDataSet(req *rda.DBRequest) (*rda.DataSet, error) {

	connectionString, _ := buildConnectionString(req)
	conMgr := GetConnectionManager()
	db, err := conMgr.GetConnection(connectionString)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(req.SqlStatement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	dataSet := &rda.DataSet{
		Tables: []*rda.DataTable{},
	}

	query, err := stmt.Query()
	if err != nil {
		return nil, err		
	}

	for {
		table, err := createTable(query) 
		for query.Next() {
			// Since the query.Scan(dest ...interface{}) takes
			// a slice of pointer, we need to create two slice
			// one for actual values, and one for the pointer to
			// each actual values. Just pass the pointer slice
			// to scan method to make things work.			
			values, valuePtrs := table.InitValueSlots()
			err = query.Scan(valuePtrs...)
			if err != nil {
				log.Println("Scan failed:", err)
				return nil, err
			}

			table.AddRow(values)
		}

		// Add Current table into this data set
		dataSet.AddTable(table)

		// If no more result set found in this query
		// finish execution
		if !query.NextResultSet() {
			break
		}
	}
	return dataSet, nil	
}
 
func createTable(query *sql.Rows) (*rda.DataTable, error) {
	columnTypes, _ := query.ColumnTypes()

	table := &rda.DataTable{
		Columns: make([]*rda.DataColumn, len(columnTypes)),
		Rows:    make([]*rda.DataRow, 0, 10),
	}

	for i:=0;i<len(columnTypes);i++{ 
		columnType := columnTypes[i]
		
		column := &rda.DataColumn{
			Index : int32(i),
			Name : columnType.Name(),
			DbType: columnType.DatabaseTypeName(),		
			Type: columnType.ScanType().String(),
		}

		if length, ok := columnType.Length(); ok{
			column.Length = length
		}

		if precision, scale, ok := columnType.DecimalSize(); ok{
			column.Precision = precision
			column.Scale = scale
		}

		table.Columns[i] = column
	}
 
	return table, nil
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
