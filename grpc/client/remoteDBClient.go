package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/xiaoyumu/go-study/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var group sync.WaitGroup

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(3600*time.Second))

	if err != nil {
		log.Fatalf("Failed to connect to remove server: %v", err)
	}
	defer conn.Close()
	c := proto.NewRemoteDBServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	//sqlStatementForScalar := "SELECT 123.99, SYSDATETIMEOFFSET()"

	sqlStatementDataSet :=
		`SELECT GETDATE() AS DateTimeColumn, 
		SYSDATETIMEOFFSET() AS DateTimeOffsetColumn, 
		CAST(255 AS TINYINT) AS TinyIntValueColumn, 
		CAST(100 AS INT) AS IntValueColumn, 
		CAST(200 AS BIGINT) AS BigIntValueColumn, 
		CAST(10.99 AS Decimal(18,4)) AS DecimalColumn,
		CAST(9.99999999 AS Decimal(18,8)) AS Decimal2Column,
		CAST(1 AS BIT) AS BooleanColumn,
		CAST(N'<Test/>' AS XML) AS XmlColumn,
		CAST('x' AS CHAR) AS SingleChar,		
		CAST(N'这是一个中文字符串' AS NVARCHAR(20)) AS UnicodeString,
		NEWID() AS UUID,
		null AS ValueNull 
	 SELECT 'Hahaha' AS TextColumn`

	//group.Add(2)

	//tryExecuteNoneQuery(ctx, c)
	tryExecuteDataSet(ctx, c, sqlStatementDataSet)
	/*
		go func() {
			for x := 0; x < 100; x++ {
				tryExecuteScalar(ctx, c, sqlStatementForScalar)
				time.Sleep(300 * time.Millisecond)
			}
			group.Done()
		}()

		go func() {
			for x := 0; x < 100; x++ {
				tryExecuteDataSet(ctx, c, sqlStatementDataSet)
				time.Sleep(330 * time.Millisecond)
			}
			group.Done()
		}()
	*/

	//group.Wait()
	log.Println("Done")
}

func getServerInfo() *proto.ServerInfo {
	return &proto.ServerInfo{
		Server:   "192.168.1.154",
		Port:     1433,
		UserId:   "dev",
		Password: "d3v",
		Database: "godemo",
	}
}

func logFatal(err error) {
	if err != nil {
		log.Fatalf("Faild to call remote DB service : %v", err)
	}
}

func dumpRemoteResponse(response *proto.DBResponse) {
	log.Printf("Remote Result is : %v", response.Succeeded)
	log.Printf("Remote Message is : %s", response.Message)
	log.Printf("Remote ScalarValue is : %v", response.ScalarValue)
	log.Printf("Remote Dataset is : %v", response.Dataset)
	log.Printf("Remote RowEffected is : %v", response.RowEffected)
}

func tryExecuteScalar(ctx context.Context, client proto.RemoteDBServiceClient, sqlStatement string) {
	response, err := client.ExecuteScalar(ctx, &proto.DBRequest{
		ServerInfo:   getServerInfo(),
		SqlStatement: sqlStatement,
	})
	logFatal(err)

	dumpRemoteResponse(response)

	if response.ScalarValue == nil {
		log.Println("ScalerValue is nil in the response.")
		return
	}

	value := proto.DecodeValue(
		response.ScalarValue.DbType,
		response.ScalarValue.Type,
		response.ScalarValue.Value)

	log.Printf("Scalar Value is [%v]", value)
}

func tryExecuteDataSet(ctx context.Context, client proto.RemoteDBServiceClient, sqlStatement string) {
	log.Println("--------------------------------------------------------------")

	log.Printf("Executing SQL: ")
	log.Println(sqlStatement)

	start := time.Now()
	response, err := client.ExecuteDataSet(ctx, &proto.DBRequest{
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
	dumpDataSet(response.Dataset)
}

func dumpDataSet(ds *proto.DataSet) {
	tables := ds.GetTables()
	log.Printf("Total %d tables in the returned data set.", len(tables))

	for i, table := range tables {
		log.Printf("(Table[%d]) %s :", i, table.GetName())
		log.Printf("  %6s  %-21s %-16s %-10s %-11s %-11s %s",
			"Column",
			"Name",
			"DBType",
			"Precision",
			"Length",
			"Type",
			"Nullable")

		for c, col := range table.Columns {
			precision := ""
			if col.GetDbType() == "DECIMAL" {
				precision = fmt.Sprintf("(%d,%d)", col.GetPrecision(), col.GetScale())
			}
			log.Printf("  %6d  %-21s %-16s %-10s %-11d %-11s %v",
				c,
				col.GetName(),
				col.GetDbType(),
				precision,
				col.GetLength(),
				col.GetType(),
				col.GetNullable())
		}

		for j, row := range table.GetRows() {
			log.Printf("  (Row[%d]):", j)
			for k, cell := range row.GetValues() {
				column := table.Columns[k]
				rawValue := cell.GetValue() // []byte
				valueType := column.GetType()
				decodedValue := table.Decode(cell)
				log.Printf("    Cell[%d]:\tDBType: [%s] \tType: [%s] \tDecoded Value: [%v] RawValue:%v",
					k,
					column.DbType,
					valueType,
					decodedValue,
					rawValue)
			}
		}
	}
}

func tryExecuteNoneQuery(ctx context.Context, client proto.RemoteDBServiceClient) {
	request := &proto.DBRequest{
		ServerInfo: getServerInfo(),
		SqlStatement: `INSERT INTO dbo.Product(Name, Description, Price, CreatedTime, Quantity,UpdatedTime, Active)	
					  VALUES(@Name, @Description, @Price, @CreatedTime, @Quantity, @UpdatedTime, @Active)`,
		Parameters: make([]*proto.DBParameter, 7),
	}

	request.Parameters[0] = (&proto.DBParameter{Name: "@Name"}).SetDBType("VARCHAR").SetLength(50).SetValue("Test Name")
	request.Parameters[1] = (&proto.DBParameter{Name: "@Description"}).SetDBType("NVARCHAR").SetLength(500).SetValue("Test Description")
	request.Parameters[2] = (&proto.DBParameter{Name: "@Price"}).SetDBType("DECIMAL").SetValue(1258.99)
	request.Parameters[3] = (&proto.DBParameter{Name: "@CreatedTime"}).SetDBType("DATETIME").SetValue(time.Now())
	request.Parameters[4] = (&proto.DBParameter{Name: "@Quantity"}).SetDBType("INT").SetValue(12)
	request.Parameters[5] = (&proto.DBParameter{Name: "@UpdatedTime"}).SetDBType("DATETIME").SetValue(time.Now())
	request.Parameters[6] = (&proto.DBParameter{Name: "@Active"}).SetDBType("BIT").SetValue(true)

	response, err := client.ExecuteNoneQuery(ctx, request)
	logFatal(err)

	dumpRemoteResponse(response)
}
