package common

import (
	"fmt"
	"strconv"
	"time"
)

// Int16 function returns a pointor to a int16 value
func Int16() *int16 {
	var value int16
	return &value
}

// Int32 function returns a pointor to a int32 value
func Int32() *int32 {
	var value int32
	return &value
}

// Int64 function returns a pointor to a int64 value
func Int64() *int64 {
	var value int64
	return &value
}

// Float32 function returns a pointor to a float32 value
func Float32() *float32 {
	var value float32
	return &value
}

// Float64 function returns a pointor to a float64 value
func Float64() *float64 {
	var value float64
	return &value
}

// String function returns a pointor to a string value
func String() *string {
	var value string
	return &value
}

// Bool function returns a pointor to a bool value
func Bool() *bool {
	var value bool
	return &value
}

// GetTypedObjectPtr function Return pointer of typed object according to the given typeString
func GetTypedObjectPtr(typeString string) (interface{}, error) {
	switch typeString {
	case "time.Time":
		return &time.Time{}, nil
	case "Time":
		return &time.Time{}, nil
	case "float32":
		return Float32(), nil
	case "float64":
		return Float64(), nil
	case "int16":
		return Int16(), nil
	case "int32":
		return Int32(), nil
	case "int64":
		return Int64(), nil
	case "string":
		return String(), nil
	case "bool":
		return Bool(), nil
	}

	return nil, fmt.Errorf("unable to determine the type by [%s]", typeString)
}

// ToFloat64 function converts []byte slice to float64 value
func ToFloat64(value []byte) (float64, error) {
	floatValue, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}
