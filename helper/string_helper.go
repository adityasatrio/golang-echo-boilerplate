package helper

import (
	"encoding/json"
	"errors"
	"fmt"
)

func StringToJson(inputStr string) map[string]interface{} {
	var jsonOutput map[string]interface{}
	err := json.Unmarshal([]byte(inputStr), &jsonOutput)
	if err != nil {
		errStr := fmt.Errorf("error unmarshalling JSON: %v", err)
		errors.New(errStr.Error())
	}

	return jsonOutput
}
