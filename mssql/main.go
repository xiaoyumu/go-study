package main

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	connectionString := "sqlserver://dev:d3v@192.168.1.154:1433?database=godemo&connection+timeout=30"
	log.Printf("Connecting to %s", connectionString)
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	log.Println("Connected")
	log.Printf("Sending ping to SQL Server ...")
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	log.Println("Ping succeeded")
	defer db.Close()

	stmt, err := db.Prepare(
		"SELECT ID, Name, Description, " +
			"Price, " +
			"CAST(Price * 100 AS BigInt) AS PriceInt, CreatedTime, Quantity, UpdatedTime, Active " +
			"FROM dbo.Product WITH(NOLOCK)")

	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}

	defer stmt.Close()

	query, err := stmt.Query()

	var id int32
	var name string
	var description string
	var price []byte
	var priceInt int64
	var createdTime time.Time
	var quantity int32
	var updatedTime time.Time
	var active bool

	for {
		if !query.Next() {
			break
		}
		err = query.Scan(&id, &name, &description, &price, &priceInt, &createdTime, &quantity, &updatedTime, &active)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		log.Printf("Row: %v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t",
			id, name, description, toFloat64(price), priceInt, createdTime, quantity, updatedTime, active)
	}

}

func toFloat64(value []uint8) float64 {
	floatValue, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		panic(err)
	}
	return floatValue
}

func float64frombytes(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
