package proto

// AddRow function builds a DataRow which contains all the value in the parameter rowValues
func (table *DataTable) AddRow(rowValues []interface{}) {
	row := &DataRow{}
	//valueCount := len(rowValues)
	for idx, value := range rowValues {
		collectValueInRow(row, idx, value)
	}
	table.Rows = append(table.Rows, row)
}

func collectValueInRow(row *DataRow, index int, value interface{}) {
	tryCollectString(row, index, value)
}

func tryCollectString(row *DataRow, index int, value interface{}) {

}
