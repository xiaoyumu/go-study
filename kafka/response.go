package main

import "encoding/json"

const (
	// TraceResponseMessageType defines message type id.
	TraceResponseMessageType int32 = 9528
)

// TraceResponseMessage defines message format.
type TraceResponseMessage struct {
	Message string `json:"Message"`
	Status  string `json:"Status"`
}

// MessageNumber returns the message number.
func (msg TraceResponseMessage) MessageNumber() int32 {
	return TraceResponseMessageType
}

// Serialize serializes Message into bytes.
func (msg TraceResponseMessage) Serialize() ([]byte, error) {
	return json.Marshal(msg)
}
