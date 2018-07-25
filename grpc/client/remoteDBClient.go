package main

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/ptypes/timestamp"

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
		log.Fatalf("Failed to connect to remove server: %v", err)
	}
	defer conn.Close()
	c := rda.NewRemoteDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tryExecuteDataSet(ctx, c)

}

func getServerInfo() *rda.ServerInfo {
	return &rda.ServerInfo{
		Server:   "192.168.1.154",
		Port:     1433,
		UserId:   "dev",
		Password: "d3v",
		Database: "godemo",
	}
}

func tryExecuteDataSet(ctx context.Context, client rda.RemoteDBServiceClient) {
	response, err := client.ExecuteDataSet(ctx, &rda.DbRequest{
		ServerInfo:   getServerInfo(),
		SqlStatement: "SELECT GETDATE(), 1 AS Value, null AS ValueNull",
	})
	logFatal(err)

	if response.Succeeded {
		log.Println("Remote call succeeded.")
	} else {
		log.Printf("Remote call failed due to %s.", response.Message)
	}

	ds := response.Dataset

	tables := ds.GetTables()
	log.Printf("Total %d tables in the returned data set.", len(tables))

	for i, table := range tables {
		log.Printf("Dumping table %s index %d", table.GetName(), i)
		for j, row := range table.GetRows() {
			log.Printf("  Dumping Row[%d]", j)
			for k, cell := range row.GetValues() {
				column := table.Columns[k]
				valueObject := cell.GetValue() // bytes
				valueType := column.GetType()
				log.Printf("    Cell[%d]: DataType: %s Value: %v", k, valueType, valueObject)
			}
		}
	}

}

func logFatal(err error) {
	if err != nil {
		log.Fatalf("Faild to call remote DB service : %v", err)
	}
}

func dumpRemoteResponse(response *rda.DbResponse) {
	log.Printf("Remote Result is : %v", response.Succeeded)
	log.Printf("Remote Message is : %s", response.Message)
	log.Printf("Remote ScalarValue is : %v", response.ScalarValue)
	log.Printf("Remote Dataset is : %v", response.Dataset)
}

func tryExecuteScalar(ctx context.Context, client rda.RemoteDBServiceClient) {
	response, err := client.ExecuteScalar(ctx, &rda.DbRequest{
		ServerInfo:   getServerInfo(),
		SqlStatement: "SELECT GETDATE(), 1",
	})
	logFatal(err)

	dumpRemoteResponse(response)

	if response.ScalarValue == nil {
		log.Println("ScalerValue is nil in the response.")
		return
	}

	var dynamicValue ptypes.DynamicAny

	// The second parameter of ptypes.UnmarshalAny() method should be an address
	// of ptypes.DynamicAny, so & must be provided. Otherwise, an error will be
	// throw with message:
	//    mismatched message type: got "google.protobuf.Timestamp" want ""

	/*
		if err := ptypes.UnmarshalAny(response.ScalarValue.Value, &dynamicValue); err != nil {
			log.Println("Failed to unmarshal Any due to " + err.Error())
			os.Exit(-1)
		}*/

	if ts, ok := dynamicValue.Message.(*timestamp.Timestamp); ok {
		time, _ := ptypes.Timestamp(ts)
		log.Println(time)
	}
}

func tryExecuteNoneQuery(ctx context.Context, client rda.RemoteDBServiceClient) {
	response, err := client.ExecuteNoneQuery(ctx, &rda.DbRequest{
		ServerInfo:   getServerInfo(),
		SqlStatement: "SELECT GETDATE(), 1",
	})
	logFatal(err)

	dumpRemoteResponse(response)

	if response.ScalarValue == nil {
		log.Println("ScalerValue is nil in the response.")
		return
	}

	var dynamicValue ptypes.DynamicAny

	// The second parameter of ptypes.UnmarshalAny() method should be an address
	// of ptypes.DynamicAny, so & must be provided. Otherwise, an error will be
	// throw with message:
	//    mismatched message type: got "google.protobuf.Timestamp" want ""
	/*
		if err := ptypes.UnmarshalAny(response.ScalarValue.Value, &dynamicValue); err != nil {
			log.Println("Failed to unmarshal Any due to " + err.Error())
			os.Exit(-1)
		}*/

	if ts, ok := dynamicValue.Message.(*timestamp.Timestamp); ok {
		time, _ := ptypes.Timestamp(ts)
		log.Println(time)
	}
}
