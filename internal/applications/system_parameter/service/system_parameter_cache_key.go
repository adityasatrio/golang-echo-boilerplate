package service

import (
	"myapp/internal/vars"
	"strconv"
	"strings"
)

const (
	Separator       = ":"
	SystemParameter = "system-parameter"
)

func CacheKeySysParamWithId(id int) string {
	return strings.Join([]string{vars.ApplicationName(), SystemParameter, strconv.Itoa(id)}, Separator)
}

func CacheKeySysParams() string {
	return strings.Join([]string{vars.ApplicationName(), SystemParameter, "list"}, Separator)
}
