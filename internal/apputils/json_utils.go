package apputils

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
)

func GetFieldBytes(json []byte, path string) (interface{}, error) {
	result := gjson.GetBytes(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return nil, errors.New(errMsg)
	}

	return result.Value(), nil
}

func GetField(json string, path string) (interface{}, error) {
	result := gjson.Get(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return nil, errors.New(errMsg)
	}

	return result.Value(), nil
}

func GetResultBytes(json []byte, path string) (gjson.Result, error) {
	result := gjson.GetBytes(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return result, errors.New(errMsg)
	}

	return result, nil
}

func GetResult(json string, path string) (gjson.Result, error) {
	result := gjson.Get(json, path)
	if !result.Exists() {
		errMsg := fmt.Sprintf("path not exist : %s # raw json is : %s", path, string(json))
		log.Error(errMsg)

		return result, errors.New(errMsg)
	}

	return result, nil
}
