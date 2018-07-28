package proto

import (
	"fmt"

	"github.com/Kelindar/binary"
)

// CreateDataSet function returns new instance of DataSet
func CreateDataSet() *DataSet {
	return &DataSet{
		Tables: make([]*DataTable, 0, 1),
	}
}

// AddTable function adds a table into the DataSet
func (ds *DataSet) AddTable(table *DataTable) {
	if len(table.Name) == 0 {
		table.Name = fmt.Sprintf("Table_%v", len(ds.Tables)+1)
	}
	ds.Tables = append(ds.Tables, table)
}

// AddRow function builds a DataRow which contains all the value in the parameter rowValues
func (table *DataTable) AddRow(rowValues []interface{}) {
	row := &DataRow{
		Values: make([]*DBValue, len(rowValues)),
	}
	for idx, value := range rowValues {
		row.Values[idx] = &DBValue{
			Index: int32(idx),
			Value: Serialize(value),
		}
	}
	table.Rows = append(table.Rows, row)
}

// Serialize convert the value into a slice of bytes
func Serialize(value interface{}) []byte {
	if value == nil {
		return nil
	}
	bytes, err := binary.Marshal(value)
	if err != nil {
		return nil
	}
	return bytes
}

// InitValueSlots function initialize a pair of value slices for sql scanning
func (table *DataTable) InitValueSlots() ([]interface{}, []interface{}) {
	return CreateValueSlotForScan(len(table.Columns))
}

// CreateValueSlotForScan function create and initialize a pair of value slices for sql row scanning
func CreateValueSlotForScan(columnCount int) ([]interface{}, []interface{}) {
	values := make([]interface{}, columnCount)
	valuePtrs := make([]interface{}, columnCount)
	// Store the address of each value in values slice into
	// corresponding element of valuePtrs slice
	for i := 0; i < columnCount; i++ {
		valuePtrs[i] = &values[i]
	}

	return values, valuePtrs
}
