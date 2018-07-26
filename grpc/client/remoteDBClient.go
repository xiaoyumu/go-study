package main

import (
	"github.com/Kelindar/binary"
	"log"
	"time"
	"fmt"
	"reflect"
	"strconv"
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
	log.Println("--------------------------------------------------------------")
	sqlStatement := 
	`SELECT GETDATE() AS DateTimeColumn, 
		CAST(255 AS TINYINT) AS TinyIntValueColumn, 
		CAST(100 AS INT) AS IntValueColumn, 
		CAST(200 AS BIGINT) AS BigIntValueColumn, 
		CAST(10.99 AS Decimal(18,4)) AS DecimalColumn,
		null AS ValueNull 
	 SELECT 'Hahaha' AS TextColumn`
					
	log.Printf("Executing SQL: ")
	log.Println(sqlStatement)


	start := time.Now()
	response, err := client.ExecuteDataSet(ctx, &rda.DBRequest{
		ServerInfo:   getServerInfo(),
		SqlStatement: sqlStatement,
	})
	elapsed := time.Since(start)
    log.Printf("ExecuteDataSet took %s", elapsed)
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
		log.Printf("(Table[%d]) %s :", i, table.GetName())
		for c, col := range table.GetColumns(){
			log.Printf("  Column[%d] \tName: [%s] \tDBType: [%s](%d,%d) \tLTH: [%d] \tNullable: [%v]", 
			c, 
			col.GetName(),
			col.GetDbType(),
			col.GetPrecision(),
			col.GetScale(),
			col.GetLength(),
			col.GetNullable())
		}

		for j, row := range table.GetRows() {
			log.Printf("  (Row[%d]):", j)
			for k, cell := range row.GetValues() {
				column := table.Columns[k]
				valueObject := cell.GetValue() // []byte
				valueType := column.GetType()
				log.Printf("    Cell[%d]: \tType: [%s] \tValue: [%v]", k, valueType, decodeValue(column, valueObject))
			}
		}
	}
}

func decodeValue(column *rda.DataColumn, value []byte) interface{} {
	if value == nil{
		return "[NULL]"
	}
	if len(column.Type) == 0{
		return value
	}

	if column.GetDbType() == "DECIMAL" {
		decodedDecimalRaw := make([]uint8, 8)
		err := binary.Unmarshal(value, &decodedDecimalRaw)
		if err != nil {
			log.Printf("Faile to Unmarshal value %v into []uint8 due to %s", value, err)
			return value
		}
		if float64Value, err := toFloat64(decodedDecimalRaw); err != nil{
			log.Printf("Faile to convert value %v into float64 due to %s", value, err)
			return value
		} else {
			return float64Value
		}
	}

	originalType, err := getOriginalType(column.Type)
	if err != nil {
		return value
	}
	decodedValue := reflect.New(originalType)
	errDecode := binary.Unmarshal(value, decodedValue.Interface())
	if errDecode != nil {
		return value
	}
	return decodedValue.Elem()
}

func toFloat64(value []byte) (float64, error) {
	floatValue, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}

func getOriginalType(typeString string) (reflect.Type, error){
	switch typeString{
		case "time.Time":
			return reflect.TypeOf((*time.Time)(nil)).Elem(), nil
		case "float32":
			return reflect.TypeOf((*float32)(nil)).Elem(), nil
		case "int32":
			return reflect.TypeOf((*int32)(nil)).Elem(), nil
		case "int64":
			return reflect.TypeOf((*int64)(nil)).Elem(), nil
		case "string":
			return reflect.TypeOf((*string)(nil)).Elem(), nil
	}

	return nil, fmt.Errorf("unable to determine the type by [%s]", typeString)
}

func logFatal(err error) {
	if err != nil {
		log.Fatalf("Faild to call remote DB service : %v", err)
	}
}

func dumpRemoteResponse(response *rda.DBResponse) {
	log.Printf("Remote Result is : %v", response.Succeeded)
	log.Printf("Remote Message is : %s", response.Message)
	log.Printf("Remote ScalarValue is : %v", response.ScalarValue)
	log.Printf("Remote Dataset is : %v", response.Dataset)
}

func tryExecuteScalar(ctx context.Context, client rda.RemoteDBServiceClient) {
	response, err := client.ExecuteScalar(ctx, &rda.DBRequest{
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
	response, err := client.ExecuteNoneQuery(ctx, &rda.DBRequest{
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
