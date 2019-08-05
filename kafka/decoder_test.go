package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonString := `{ "Key": {"Type": "Operation"}, "Value": {"Name": "Name1"}}`
	traceMsg, err := Decode([]byte(jsonString))
	assert.Nil(t, err)
	assert.NotNil(t, traceMsg)
	assert.NotNil(t, traceMsg.Key)
	assert.NotNil(t, traceMsg.Value)
}
