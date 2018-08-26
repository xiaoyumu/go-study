package proto

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	pt "github.com/golang/protobuf/ptypes"
	du "github.com/golang/protobuf/ptypes/duration"
	ts "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ToProtoBufBytes converts value to []byte
func ToProtoBufBytes(value interface{}) []byte {
	if value == nil {
		return nil
	}

	t := reflect.TypeOf(value)
	isAPointer := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isAPointer = true
	}
	bytesRef := make([]byte, 0, 0)
	deterministic := true
	typeDetermined := false
	var bytes []byte
	var marshalErr error

	if v, ok := value.(time.Duration); ok {
		typeDetermined = true
		bytes, marshalErr = durationToProtoBufBytes(v, bytesRef, deterministic)
	}

	if !typeDetermined {
		if vp, ok := value.(*time.Duration); ok {
			typeDetermined = true
			bytes, marshalErr = durationToProtoBufBytes(*vp, bytesRef, deterministic)
		}
	}

	if !typeDetermined {
		if v, ok := value.(time.Time); ok {
			typeDetermined = true
			bytes, marshalErr = timeToProtoBufBytes(v, bytesRef, deterministic)
		}
	}

	if !typeDetermined {
		if vp, ok := value.(*time.Time); ok {
			typeDetermined = true
			bytes, marshalErr = timeToProtoBufBytes(*vp, bytesRef, deterministic)
		}
	}

	if typeDetermined {
		logSerializeError(value, marshalErr)
		return bytes
	}

	switch t.Kind() {
	case reflect.String:
		typeDetermined = true
		v, _ := value.(string)
		if isAPointer {
			if vp, _ := value.(*string); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.StringValue{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		typeDetermined = true
		v, _ := value.(int32)
		if isAPointer {
			if vp, _ := value.(*int32); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.Int32Value{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Int64:
		typeDetermined = true
		v, _ := value.(int64)
		if isAPointer {
			if vp, _ := value.(*int64); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.Int64Value{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		typeDetermined = true
		v, _ := value.(uint32)
		if isAPointer {
			if vp, _ := value.(*uint32); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.UInt32Value{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Uint64:
		typeDetermined = true
		v, _ := value.(uint64)
		if isAPointer {
			if vp, _ := value.(*uint64); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.UInt64Value{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Float32:
		typeDetermined = true
		v, _ := value.(float32)
		if isAPointer {
			if vp, _ := value.(*float32); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.FloatValue{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Float64:
		typeDetermined = true
		v, _ := value.(float64)
		if isAPointer {
			if vp, _ := value.(*float64); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.DoubleValue{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break

	case reflect.Bool:
		typeDetermined = true
		v, _ := value.(bool)
		if isAPointer {
			if vp, _ := value.(*bool); vp != nil {
				v = *vp
			}
		}
		bytes, marshalErr = (&wrappers.BoolValue{Value: v}).XXX_Marshal(bytesRef, deterministic)
		break
	}

	logSerializeError(value, marshalErr)

	return bytes
}

func logSerializeError(value interface{}, err error) {
	if err == nil {
		return
	}

	log.Printf("Failed serialize value %v into protobuf bytes ([]byte) due to %s",
		value,
		err)
}

func durationToProtoBufBytes(value time.Duration, bytesRef []byte, deterministic bool) ([]byte, error) {
	d := pt.DurationProto(value)
	return d.XXX_Marshal(bytesRef, deterministic)
}

func timeToProtoBufBytes(value time.Time, bytesRef []byte, deterministic bool) ([]byte, error) {
	ts, convErr := pt.TimestampProto(value)

	if convErr == nil {
		return ts.XXX_Marshal(bytesRef, deterministic)
	}

	log.Printf("Failed to convert value %v into TimestampProto due to %s",
		value,
		convErr.Error())
	return nil, convErr
}

// FromProtoBufBytes converts the proto buf data bytes back to value
func FromProtoBufBytes(typeName string, data []byte) (interface{}, error) {
	switch strings.ToLower(typeName) {
	case "string", "char", "nchar", "varchar", "nvarchar", "text", "ntext":
		v := &wrappers.StringValue{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "bool", "boolean", "bit":
		v := &wrappers.BoolValue{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "int", "int8", "int16", "int32", "tinyint", "samllint":
		v := &wrappers.Int32Value{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "int64", "long", "bigint":
		v := &wrappers.Int64Value{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "uint", "uint8", "uint16", "uint32":
		v := &wrappers.UInt32Value{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "uint64":
		v := &wrappers.UInt64Value{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "float32", "numeric":
		v := &wrappers.FloatValue{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "double", "decimal", "float64", "money", "real":
		v := &wrappers.DoubleValue{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return v.Value, nil

	case "time":
		v := &du.Duration{}
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		return pt.Duration(v)

	case "time.time", "date", "datetime", "datetime2", "datetimeoffset":
		v := &ts.Timestamp{} // Unmarshal bytes to proto Timestamp
		if err := v.XXX_Unmarshal(data); err != nil {
			return nil, err
		}
		// Then try to convert it back to time.Time
		// No timezone information currently
		return pt.Timestamp(v)
	}

	return nil, fmt.Errorf("unsupported type [%s]", typeName)
}
