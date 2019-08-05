package main

import "encoding/json"

const (
	MessageTypeID int32 = 9527
)


type TraceMessage struct {
	Key  *json.RawMessage
	Value *json.RawMessage
}

type TraceResponse struct {
	Status string
	Message string
}

func (resp TraceResponse) Serialize() []byte {
	data, err:= json.Marshal(resp)
	if err != nil{
		return nil
	}
	return data
}
