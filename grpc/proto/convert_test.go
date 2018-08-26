package proto

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ToProtoBufBytes_String(t *testing.T) {
	bytes := ToProtoBufBytes(" #$%^ 123 English Español 中文 བོད་སྐད་ ภาษาไทย にっぽんご 한글 العَرَبِيَّة ")
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("string", bytes)
	assert.Nil(t, err)
	assert.Equal(t, " #$%^ 123 English Español 中文 བོད་སྐད་ ภาษาไทย にっぽんご 한글 العَرَبِيَّة ", value)
}

func Test_ToProtoBufBytes_String2(t *testing.T) {
	bytes := ToProtoBufBytes("29051")
	assert.NotNil(t, bytes)
	t.Logf("%v Base64: %s", bytes, base64.StdEncoding.EncodeToString(bytes))

	value, err := FromProtoBufBytes("string", bytes)
	assert.Nil(t, err)
	assert.Equal(t, "29051", value)
}

func Test_ToProtoBufBytes_StringPointer(t *testing.T) {
	stringValue := "29051"
	bytes := ToProtoBufBytes(&stringValue)
	assert.NotNil(t, bytes)
	t.Logf("%v Base64: %s", bytes, base64.StdEncoding.EncodeToString(bytes))

	value, err := FromProtoBufBytes("string", bytes)
	assert.Nil(t, err)
	assert.Equal(t, "29051", value)
}

func Test_ToProtoBufBytes_Int32(t *testing.T) {
	var originalValue int32
	originalValue = 123456789
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("int32", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
}

func Test_ToProtoBufBytes_Int32Pointer(t *testing.T) {
	var originalValue int32
	originalValue = 123456789
	bytes := ToProtoBufBytes(&originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("int32", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
}

func Test_ToProtoBufBytes_Int64(t *testing.T) {
	var originalValue int64
	originalValue = 99999999999
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("Int64", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
}

func Test_ToProtoBufBytes_UInt32(t *testing.T) {
	var originalValue uint32
	originalValue = 9999
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("uint32", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
}

func Test_ToProtoBufBytes_UInt64(t *testing.T) {
	var originalValue uint64
	originalValue = 99999999999
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("uint64", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}

func Test_ToProtoBufBytes_Float32(t *testing.T) {
	var originalValue float32
	originalValue = 99999999999
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("float32", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}

func Test_ToProtoBufBytes_Float64(t *testing.T) {
	var originalValue float64
	originalValue = 99999999.9999999
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("float64", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}

func Test_ToProtoBufBytes_Time(t *testing.T) {
	originalValue := time.Date(2018, 8, 25, 15, 31, 00, 999987237, time.UTC)

	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("time.Time", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}

func Test_ToProtoBufBytes_TimePointer(t *testing.T) {
	originalValue := time.Date(2018, 8, 25, 15, 31, 00, 999987237, time.UTC)

	bytes := ToProtoBufBytes(&originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("time.Time", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}

func Test_ToProtoBufBytes_Duration(t *testing.T) {
	originalValue, _ := time.ParseDuration("1s")
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("time", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}

func Test_ToProtoBufBytes_Bool(t *testing.T) {
	var originalValue bool
	originalValue = true
	bytes := ToProtoBufBytes(originalValue)
	assert.NotNil(t, bytes)
	value, err := FromProtoBufBytes("bool", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
	originalValue = false
	bytes = ToProtoBufBytes(originalValue)
	value, err = FromProtoBufBytes("bool", bytes)
	assert.Nil(t, err)
	assert.Equal(t, originalValue, value)
	assert.IsType(t, originalValue, value)
}
