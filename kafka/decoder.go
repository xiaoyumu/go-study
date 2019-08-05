package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Decode the given data into Message of return error
func Decode(data []byte) (*TraceMessage, error){
	if data == nil {
		return nil, errors.New("input data was nil, unable to deserialize it to TraceMessage")
	}
	// The data should be in json format, so we
	// unmarshal it with json
	var traceMsg TraceMessage
	if err := json.Unmarshal(data, &traceMsg); err != nil{
		return nil, fmt.Errorf("failed to unmarshal the data into json map due to %s", err)
	}

	if traceMsg.Key == nil{
		return nil, errors.New("invalid trace message, Key cannot be empty")
	}

	if traceMsg.Value == nil{
		return nil, errors.New("invalid trace message, Value cannot be empty")
	}

	return &traceMsg, nil
}
