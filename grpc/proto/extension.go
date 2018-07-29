package proto

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/Kelindar/binary"
	"github.com/xiaoyumu/go-study/grpc/common"
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

// Decode a DBValue
func (table *DataTable) Decode(value *DBValue) interface{} {
	column := table.Columns[value.Index]
	return DecodeValue(column.DbType, column.Type, value.Value)
}

// DeserializeValue convert DBParameter.Value back to the typed value in interface{}
func (param *DBParameter) DeserializeValue() interface{} {
	value := DecodeValue(param.DbType, param.Type, param.Value)
	return value
}

// SetValue function serialize the value parameter save it into DBParameter.Value
func (param *DBParameter) SetValue(value interface{}) *DBParameter {
	param.Value = Serialize(value)
	if value != nil {
		param.Type = reflect.TypeOf(value).Name()
	}
	return param
}

// SetDBType function sets the dbType into DBParameter.DbType
func (param *DBParameter) SetDBType(dbType string) *DBParameter {
	param.DbType = dbType
	return param
}

// SetLength function sets the dbLength into DBParameter.DbLength
func (param *DBParameter) SetLength(dbLength int64) *DBParameter {
	param.DbLength = dbLength
	return param
}

// SetType function sets the type into DBParameter.Type
func (param *DBParameter) SetType(typeString string) *DBParameter {
	param.Type = typeString
	return param
}

func decodeDecimal(valueType string, value []byte) interface{} {
	// Try to unmarshal it into float64 at the first
	if valueType == "float64" {
		float64Value := common.Float64()
		unmarshalErr := binary.Unmarshal(value, float64Value)
		if unmarshalErr == nil && float64Value != nil {
			return *float64Value
		}
	}

	decodedDecimalRaw := make([]uint8, 8)
	err := binary.Unmarshal(value, &decodedDecimalRaw)
	if err != nil {
		log.Printf("Faile to Unmarshal value %v into []uint8 due to %s", value, err)
		return value
	}

	if float64Value, err := common.ToFloat64(decodedDecimalRaw); err != nil {
		log.Printf("Faile to convert value %v into float64 due to %s", value, err)
		return value
	} else {
		return float64Value
	}
}

// DecodeValue converts binary serialized value ([]byte) back to typed value in interface{}
// according to the given dbType and valueType.
func DecodeValue(dbType string, valueType string, value []byte) interface{} {
	if value == nil {
		return nil
	}

	if dbType == "DECIMAL" {

		return decodeDecimal(valueType, value)
	}

	switch valueType {
	case "time.Time", "Time":
		val := time.Time{}
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "float32":
		var val float32
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "float64":
		var val float64
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "int16":
		var val int16
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "int32":
		var val int32
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "int64":
		var val int64
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "string":
		var val string
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	case "bool":
		var val bool
		if errDecode := binary.Unmarshal(value, &val); errDecode == nil {
			return val
		}
		return value
	}

	return value
}
