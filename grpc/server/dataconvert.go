package main

import (
	"time"

	"github.com/Kelindar/binary"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	rda "github.com/xiaoyumu/go-study/grpc/proto"
)

const defaultMessageTypePrefix string = "fish/"

// ToDBValue function wrap any value into a DBValue
func ToDBValue(index int32, value *interface{}) (*rda.DBValue, error) {
	valueBytes, err := binary.Marshal(value)
	if err != nil {
		return nil, err
	}

	return &rda.DBValue{
		Index: index,
		Value: valueBytes,
	}, nil
}

// ToAny converts DBValue to any.Any without customized prefix
func toAny(value *interface{}) (*any.Any, error) {
	return toAnyWithPrefix(value, defaultMessageTypePrefix)
}

// ToAnyWithPrefix converts DBValue to any.Any, support customized prefix
func toAnyWithPrefix(value *interface{}, messageTypePreix string) (*any.Any, error) {
	if v, ok := (*value).(time.Time); ok {
		return timeToAny(v, messageTypePreix)
	}

	if v, ok := (*value).(int64); ok {
		return int64ToAny(v, messageTypePreix)
	}

	return nil, nil
}

func int64ToAny(v int64, messageTypePreix string) (*any.Any, error) {
	/*
		if serialized, err := proto.Marshal(proto.Int64(v)); err != nil {
			return nil, err
		}*/

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
