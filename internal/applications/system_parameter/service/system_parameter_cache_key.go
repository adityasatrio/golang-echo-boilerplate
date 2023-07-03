package service

import (
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

var (
	ApplicationName = viper.GetString("application.name")
	Separator       = ":"
	SystemParameter = "system-parameter"
)

func CacheKeySysParamWithId(id int) string {
	return strings.Join([]string{ApplicationName, SystemParameter, strconv.Itoa(id)}, Separator)
}

func CacheKeySysParams() string {
	return strings.Join([]string{ApplicationName, SystemParameter, "list"}, Separator)
}
