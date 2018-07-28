package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Kelindar/binary"

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

	tryExecuteScalar(ctx, c)
	//tryExecuteDataSet(ctx, c)
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
		for c, col := range table.GetColumns() {
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
	if value == nil {
		return "[NULL]"
	}
	if len(column.Type) == 0 {
		return value
	}

	if column.GetDbType() == "DECIMAL" {
		decodedDecimalRaw := make([]uint8, 8)
		err := binary.Unmarshal(value, &decodedDecimalRaw)
		if err != nil {
			log.Printf("Faile to Unmarshal value %v into []uint8 due to %s", value, err)
			return value
		}
		if float64Value, err := toFloat64(decodedDecimalRaw); err != nil {
			log.Printf("Faile to convert value %v into float64 due to %s", value, err)
			return value
		} else {
			return float64Value
		}
	}

	originalType, err := getTypedObjectPtr(column.Type)
	if err != nil {
		return value
	}
	errDecode := binary.Unmarshal(value, originalType)
	if errDecode != nil {
		return value
	}
	return originalType
}

func toFloat64(value []byte) (float64, error) {
	floatValue, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}

func getTypedObjectPtr(typeString string) (interface{}, error) {
	switch typeString {
	case "time.Time":
		return &time.Time{}, nil
	case "float32":
		return Float32(), nil
	case "float64":
		return Float64(), nil
	case "int16":
		return Int16(), nil
	case "int32":
		return Int32(), nil
	case "int64":
		return Int64(), nil
	case "string":
		return String(), nil
	case "bool":
		return Bool(), nil
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

	value := decodeValue(response.ScalarValue.Type, response.ScalarValue.Value.Value)

	log.Printf("Scalar Value is [%v]", value)
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

}
