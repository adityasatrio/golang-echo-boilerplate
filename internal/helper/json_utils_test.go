package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFieldBytes(t *testing.T) {
	json := []byte(`{"name":"John","age":25}`)
	path := "name"

	value, err := GetFieldBytes(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", value)

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetFieldBytes(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

func TestGetField(t *testing.T) {
	json := `{"name":"John","age":25}`
	path := "name"

	value, err := GetField(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", value)

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetField(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

func TestGetResultBytes(t *testing.T) {
	json := []byte(`{"name":"John","age":25}`)
	path := "name"

	result, err := GetResultBytes(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", result.String())

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetResultBytes(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

func TestGetResult(t *testing.T) {
	json := `{"name":"John","age":25}`
	path := "name"

	result, err := GetResult(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", result.String())

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetResult(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}
