package helper

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
)

func CacheKey(entityValue interface{}, params map[string]string) string {
	// Get the name of the struct type
	t := reflect.TypeOf(entityValue)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	typeName := t.Name()

	// Get the value of the ID field of the struct
	v := reflect.ValueOf(entityValue)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	idField := v.FieldByName("ID")
	if !idField.IsValid() {
		// Handle error: struct does not have an ID field
	}
	entityID := idField.Int()

	// Encode the input parameters as a JSON string
	paramBytes, err := json.Marshal(params)
	if err != nil {
		// Handle error
	}
	paramString := string(paramBytes)

	// Generate a hash of the input parameters
	paramHash := fmt.Sprintf("%x", md5.Sum([]byte(paramString)))

	// Return the cache key
	return fmt.Sprintf("%s:%d:%s", typeName, entityID, paramHash)
}
