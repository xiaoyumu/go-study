package proto

import (
	"log"
	"time"

	"github.com/Kelindar/binary"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/xiaoyumu/go-study/grpc/common"
)

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

func decodeUniqueIdentifier(value interface{}) string {
	var uuid mssql.UniqueIdentifier

	uuid.Scan(value)

	return uuid.String()
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
