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

/*
    ExecuteNoneQuery(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error)
	ExecuteScalar(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error)
	ExecuteDataSet(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error)
*/

func (s *server) ExecuteNoneQuery(ctx context.Context, req *rda.DbRequest) (*rda.DbResponse, error) {
	log.Printf("RPC call [ExecuteNoneQuery] received From client [%s].", getClietIP(ctx))
	return &rda.DbResponse{
		Succeeded: false,
		Message:   "ExecuteNoneQuery was not implemented yet",
	}, nil
}

func (s *server) ExecuteScalar(ctx context.Context, req *rda.DbRequest) (*rda.DbResponse, error) {
	log.Printf("RPC call [ExecuteScalar] received From client [%s].", getClietIP(ctx))

	return executeScalar(req)
}

func (s *server) ExecuteDataSet(ctx context.Context, req *rda.DbRequest) (*rda.DbResponse, error) {
	log.Printf("RPC call [ExecuteDataSet] received From client [%s].", getClietIP(ctx))
	msg := ""
	result := true

	response, errE := executeDbRequest(req)

	if errE != nil {
		msg = errE.Error()
	}

	if response == nil {
		response = &rda.DbResponse{
			Succeeded: false,
			Message:   msg,
		}
	} else {
		response.Succeeded = result
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

func dumpRemoteDbRequest(req *rda.DbRequest) {
	connStr, _ := buildConnectionString(req)

	log.Printf("Dumping request : %s ...", connStr)
	log.Printf("SQL:%s", req.SqlStatement)
}

func buildConnectionString(req *rda.DbRequest) (string, error) {
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

func pingServer(db *sql.DB) error {
	log.Printf("Sending ping to SQL Server ...")
	err := db.Ping()
	if err != nil {
		log.Printf("Failed to sending ping due to: %s", err.Error())
		return err
	}
	log.Println("Ping succeeded")

	return nil
}

func openConnection(connectionString string) (*sql.DB, error) {

	log.Printf("Connecting to %s", connectionString)
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Println("Cannot connect: ", err.Error())
		return nil, err
	}
	log.Println("Connected")

	err = pingServer(db)

	log.Printf("Current open connections: %d", db.Stats().OpenConnections)

	return db, nil
}

func logAndFailOnError(prefix string, description string, err error) {
	if err != nil {
		log.Printf("[%s] %s: %s", prefix, description, err)
		panic(err)
	}
}

func executeScalar(req *rda.DbRequest) (*rda.DbResponse, error) {
	connectionString, _ := buildConnectionString(req)
	db, err := openConnection(connectionString)
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

	response := rda.DbResponse{
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

		//response.ScalarValue, _ = ptypes.MarshalAny(values[0])
	}

	return &response, nil
}

// Execute the remote Db request
func executeDbRequest(req *rda.DbRequest) (*rda.DbResponse, error) {
	connectionString, _ := buildConnectionString(req)
	db, err := openConnection(connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ds := executeDataSet(db, req.SqlStatement)

	if ds == nil {
		log.Printf("Failed to executeDataSet.")
		os.Exit(-1)
	}

	for _, table := range ds.Tables {
		log.Printf("Dumping table %s", table.Name)
		//log.Println(table.Columns)
		for _, row := range table.Rows {
			log.Println(row.Values)
		}
	}

	response := rda.DbResponse{
		Dataset: ds,
	}

	return &response, nil
}

func executeDataSet(db *sql.DB, sqlStatement string) *rda.DataSet {

	dataSet := rda.DataSet{
		Tables: []*rda.DataTable{},
	}

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal("Prepare Statement failed:", err)
		panic(err)
	}
	defer stmt.Close()

	query, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err)
		panic(err)
	}

	for {

		table, columns, err := createTable(query)
		columnCount := len(columns)
		for query.Next() {
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

			addRow(table, values)
		}

		// Add Current table into this data set
		addTable(&dataSet, table)

		// If no more result set found in this query
		// finish execution
		if !query.NextResultSet() {
			break
		}
	}
	return &dataSet
}

func addTable(ds *rda.DataSet, table *rda.DataTable) {

	if len(table.Name) == 0 {
		table.Name = fmt.Sprintf("Table_%v", len(ds.Tables)+1)
	}
	ds.Tables = append(ds.Tables, table)
}

func addRow(dt *rda.DataTable, rowValues []interface{}) {
	row := rda.DataRow{
		//ParentTable: dt,
		//Values: rowValues,
	}
	dt.Rows = append(dt.Rows, &row)
}

func createTable(query *sql.Rows) (*rda.DataTable, []string, error) {
	columns, _ := query.Columns()
	//columnTypes, _ := query.ColumnTypes()
	//colCount := len(columns)
	table := rda.DataTable{
		//Columns:     make([]string, colCount),
		//ColumnCount: colCount,
		//Rows: []DataRow{},
	}

	//table.Columns = columns

	return &table, columns, nil
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
