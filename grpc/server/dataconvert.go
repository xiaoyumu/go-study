package main

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// DBValue represent any value from DB
type DBValue struct {
	value *interface{}
}

const defaultMessageTypePrefix string = "fish/"

// ToDBValue function wrap any value  into a DBValue
func ToDBValue(value *interface{}) *DBValue {
	return &DBValue{
		value: value,
	}
}

// ToAny converts DBValue to any.Any without customized prefix
func (dbv *DBValue) ToAny() (*any.Any, error) {
	return dbv.ToAnyWithPrefix(defaultMessageTypePrefix)
}

// ToAnyWithPrefix converts DBValue to any.Any, support customized prefix
func (dbv *DBValue) ToAnyWithPrefix(messageTypePreix string) (*any.Any, error) {
	if v, ok := (*dbv.value).(time.Time); ok {
		return timeToAny(v, messageTypePreix)
	}
	return nil, nil
}

func timeToAny(v time.Time, messageTypePreix string) (*any.Any, error) {

	timeValue, err := ptypes.TimestampProto(v)
	if err != nil {
		return nil, err
	}

	serialized, err := proto.Marshal(timeValue)

	if err != nil {
		return nil, err
	}

	return &any.Any{
		TypeUrl: messageTypePreix + proto.MessageName(timeValue),
		Value:   serialized,
	}, nil
}
