package proto

import (
	"reflect"
	"strings"
)

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

// NameWithoutPrefix function returns the Name of the DBParameter without @prefix
func (param *DBParameter) NameWithoutPrefix() string {
	if strings.HasPrefix(param.Name, "@") {
		return string(([]rune(param.Name))[1:])
	}
	return param.Name
}
