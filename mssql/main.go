package main

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

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

	sqlStatement := "SELECT ID, Name, Description, " +
		"Price, " +
		"CAST(Price * 100 AS BigInt) AS PriceInt, CreatedTime, Quantity, UpdatedTime, Active " +
		"FROM dbo.Product WITH(NOLOCK)"

	ds := executeDataSet(db, sqlStatement)

	if ds == nil {
		log.Printf("Failed to executeDataSet.")
		os.Exit(-1)
	}

	for _, table := range ds.tables {
		log.Printf("Dumping table %s", table.Name)
		log.Println(table.Columns)
		for _, row := range table.Rows {
			log.Println(row.Values)
		}
	}
}

func executeDataSet(db *sql.DB, sqlStatement string) *DataSet {

	dataSet := DataSet{
		tables: []*DataTable{},
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

		table, err := createTable(query)
		for query.Next() {
			// Since the query.Scan(dest ...interface{}) takes
			// a slice of pointer, we need to create two slice
			// one for actual values, and one for the pointer to
			// each actual values. Just pass the pointer slice
			// to scan method to make things work.
			values := make([]interface{}, table.ColumnCount)
			valuePtrs := make([]interface{}, table.ColumnCount)
			// Store the address of each value in values slice into
			// corresponding element of valuePtrs slice
			for i := 0; i < table.ColumnCount; i++ {
				valuePtrs[i] = &values[i]
			}

			err = query.Scan(valuePtrs...)

			if err != nil {
				log.Fatal("Scan failed:", err)
			}

			table.addRow(values)
		}

		// Add Current table into this data set
		dataSet.addTable(table)

		// If no more result set found in this query
		// finish execution
		if !query.NextResultSet() {
			break
		}
	}
	return &dataSet
}

func (ds *DataSet) addTable(table *DataTable) {

	if len(table.Name) == 0 {
		table.Name = fmt.Sprintf("Table_%v", len(ds.tables)+1)
	}
	ds.tables = append(ds.tables, table)
}

func (dt *DataTable) addRow(rowValues []interface{}) {
	row := DataRow{
		ParentTable: dt,
		Values:      rowValues,
	}
	dt.Rows = append(dt.Rows, row)
}

func createTable(query *sql.Rows) (*DataTable, error) {
	columns, _ := query.Columns()
	//columnTypes, _ := query.ColumnTypes()
	colCount := len(columns)
	table := DataTable{
		Columns:     make([]string, colCount),
		ColumnCount: colCount,
		Rows:        []DataRow{},
	}

	table.Columns = columns

	return &table, nil
}

type DataSet struct {
	tables []*DataTable
}

type DataTable struct {
	Name        string
	Columns     []string
	Rows        []DataRow
	ColumnCount int
}

type DataColumn struct {
	Name   string
	Type   string
	DBType string
	DBSize int32
	Index  int32
}

type DataRow struct {
	ParentTable *DataTable // The table which this DataRow belongs to
	Values      []interface{}
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
