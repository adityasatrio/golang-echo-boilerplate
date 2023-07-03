package service

import (
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

var (
	ApplicationName = viper.GetString("application.name")
	Separator       = ":"
	User            = "user"
)

func CacheKeyUserWithId(id uint64) string {
	return strings.Join([]string{ApplicationName, User, strconv.FormatUint(id, 10)}, Separator)
}

func CacheKeyUsers() string {
	return strings.Join([]string{ApplicationName, User, "list"}, Separator)
}
