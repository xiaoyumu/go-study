package proto

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

// Decode a DBValue
func (table *DataTable) Decode(value *DBValue) interface{} {
	column := table.Columns[value.Index]
	return DecodeValue(column.DbType, column.Type, value.Value)
}
