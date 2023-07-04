package service

import (
	"myapp/configs"
	"strconv"
	"strings"
)

const (
	Separator = ":"
	User      = "user"
)

func CacheKeyUserWithId(id uint64) string {
	return strings.Join([]string{configs.ApplicationName, User, strconv.FormatUint(id, 10)}, Separator)
}

func CacheKeyUsers() string {
	return strings.Join([]string{configs.ApplicationName, User, "list"}, Separator)
}
