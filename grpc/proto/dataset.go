package proto

import "fmt"

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
