package globalutils

import (
	"strconv"
	"strings"
)

const (
	ApplicationName = "echo-boilerplate"
	Separator       = ":"
	SystemParameter = "system-parameter"
	User            = "user"
)

func CacheKeySysParamWithId(id int) string {
	return strings.Join([]string{ApplicationName, SystemParameter, strconv.Itoa(id)}, Separator)
}

func CacheKeySysParams() string {
	return strings.Join([]string{ApplicationName, SystemParameter, "list"}, Separator)
}

func CacheKeyUserWithId(id uint64) string {
	return strings.Join([]string{ApplicationName, User, strconv.FormatUint(id, 10)}, Separator)
}

func CacheKeyUsers() string {
	return strings.Join([]string{ApplicationName, User, "list"}, Separator)
}
