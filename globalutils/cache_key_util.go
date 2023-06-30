package globalutils

import (
	"strconv"
	"strings"
)

const (
	ApplicationName = "echo-boilerplate"
	Separator       = ":"
	SystemParameter = "system-parameter"
)

func CacheKeySysParamWithId(id int) string {
	return strings.Join([]string{ApplicationName, SystemParameter, strconv.Itoa(id)}, Separator)
}

func CacheKeySysParams() string {
	return strings.Join([]string{ApplicationName, SystemParameter, "list"}, Separator)
}
