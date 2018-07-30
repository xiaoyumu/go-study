package proto

import (
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/xiaoyumu/go-study/grpc/common"
)

// DecodeDBValueIfNeccessary decode DB value retrieved from sql server driver, and try to decode some data types
// If a data needs to be decoded, the new type will be returned along with decoded value.
func DecodeDBValueIfNeccessary(dbType string, scanType string, value interface{}) (interface{}, string, error) {
	if dbType == "DECIMAL" {
		return decodeDecimalFromDBValue(value)
	}

	if dbType == "UNIQUEIDENTIFIER" {
		return decodeUniqueIdentifierFromDBValue(value)
	}

	return value, scanType, nil
}

func decodeDecimalFromDBValue(value interface{}) (interface{}, string, error) {

	if float64Value, err := common.ToFloat64(value.([]uint8)); err != nil {
		return value, "float64", fmt.Errorf("Faile to convert value %v into float64 due to %s", value, err)
	} else {
		return float64Value, "float64", nil
	}
}

func decodeUniqueIdentifierFromDBValue(value interface{}) (interface{}, string, error) {
	var uuid mssql.UniqueIdentifier
	uuid.Scan(value)
	return uuid.String(), "string", nil
}
