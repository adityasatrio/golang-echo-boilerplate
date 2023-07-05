package service

import (
	"myapp/configs"
	"strconv"
	"strings"
)

const (
	Separator       = ":"
	SystemParameter = "system-parameter"
)

func CacheKeySysParamWithId(id int) string {
	return strings.Join([]string{configs.ApplicationName, SystemParameter, strconv.Itoa(id)}, Separator)
}

func CacheKeySysParams() string {
	return strings.Join([]string{configs.ApplicationName, SystemParameter, "list"}, Separator)
}
