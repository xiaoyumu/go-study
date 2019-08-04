package main

import (
	"encoding/json"
	"fmt"
)

const (
	// TraceMessageType defines message type id.
	TraceMessageType int32 = 9527
)

// TraceMessage defines message format.
type TraceMessage struct {
	Key   *json.RawMessage
	Value *json.RawMessage
}

// MessageNumber returns the message number.
func (msg TraceMessage) MessageNumber() int32 {
	return TraceMessageType
}

// Serialize serializes Message into bytes.
func (msg TraceMessage) Serialize() ([]byte, error) {
	return json.Marshal(msg)
}

// DeserializeTraceMessage deserializes bytes into Message.
func DeserializeTraceMessage(data []byte) (*TraceMessage, error) {
	if data == nil {
		return nil, fmt.Errorf("data is null")
	}

	var msg TraceMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
