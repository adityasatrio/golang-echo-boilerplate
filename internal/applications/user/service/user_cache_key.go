package service

import (
	"myapp/internal/vars"
	"strconv"
	"strings"
)

const (
	Separator = ":"
	User      = "user"
)

func CacheKeyUserWithId(id uint64) string {
	return strings.Join([]string{vars.ApplicationName(), User, strconv.FormatUint(id, 10)}, Separator)
}

func CacheKeyUsers() string {
	return strings.Join([]string{vars.ApplicationName(), User, "list"}, Separator)
}
