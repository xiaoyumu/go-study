package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/xiaoyumu/go-study/grpc/proto"
)

// Executor interface defines the DB operations
type Executor interface {
	ExecuteDataSet(req *proto.DBRequest) (*proto.DataSet, error)
	ExecuteNoneQuery(req *proto.DBRequest) (int32, error)
	ExecuteSalar(req *proto.DBRequest) (*proto.DBScalarValue, error)
}

type RdaExecutor struct {
	conMgr ConnectionManager
}

func NewRdaExecutor() Executor {
	return &RdaExecutor{
		conMgr: GetConnectionManager(),
	}
}

func (e *RdaExecutor) ExecuteNoneQuery(req *proto.DBRequest) (int32, error) {
	return 0, fmt.Errorf("Not implemented yet")
}

func (e *RdaExecutor) ExecuteSalar(req *proto.DBRequest) (*proto.DBScalarValue, error) {
	connectionString, _ := e.conMgr.BuildConnectionString(req)
	db, err := e.conMgr.GetConnection(connectionString)
	if err != nil {
		return nil, err
	}
	stmt, errPrepare := db.Prepare(req.SqlStatement)
	if errPrepare != nil {
		return nil, errPrepare
	}
	defer stmt.Close()

	query, errQuery := stmt.Query()
	if errQuery != nil {
		return nil, errQuery
	}

	columnTypes, errGetColumnTypes := query.ColumnTypes()
	if errGetColumnTypes != nil {
		return nil, errGetColumnTypes
	}
	sv := &proto.DBScalarValue{
		Type:  toDataColumn(columnTypes[0], 0), // Just pickup the first column
		Value: &proto.DBValue{Index: 0},
	}

	if !query.Next() {
		return nil, fmt.Errorf("no Data")
	}

	values, valuePtrs := proto.CreateValueSlotForScan(len(columnTypes))
	errScan := query.Scan(valuePtrs...)
	if errScan != nil {
		return nil, errScan
	}

	sv.Value.Value = proto.Serialize(values[0])

	return sv, nil
}

func (e *RdaExecutor) ExecuteDataSet(req *proto.DBRequest) (*proto.DataSet, error) {
	connectionString, _ := e.conMgr.BuildConnectionString(req)
	db, err := e.conMgr.GetConnection(connectionString)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(req.SqlStatement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	dataSet := &proto.DataSet{
		Tables: []*proto.DataTable{},
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

func toDataColumn(columnType *sql.ColumnType, index int) *proto.DataColumn {
	column := &proto.DataColumn{
		Index:  int32(index),
		Name:   columnType.Name(),
		DbType: columnType.DatabaseTypeName(),
		Type:   columnType.ScanType().String(),
	}

	if length, ok := columnType.Length(); ok {
		column.Length = length
	}

	if precision, scale, ok := columnType.DecimalSize(); ok {
		column.Precision = precision
		column.Scale = scale
	}
	return column
}

func createTable(query *sql.Rows) (*proto.DataTable, error) {
	columnTypes, _ := query.ColumnTypes()

	table := &proto.DataTable{
		Columns: make([]*proto.DataColumn, len(columnTypes)),
		Rows:    make([]*proto.DataRow, 0, 10),
	}

	for i := 0; i < len(columnTypes); i++ {
		columnType := columnTypes[i]
		table.Columns[i] = toDataColumn(columnType, i)
	}

	return table, nil
}
