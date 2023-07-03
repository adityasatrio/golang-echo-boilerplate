package service

import (
	"github.com/spf13/viper"
	"myapp/configs"
	"strconv"
	"strings"
)

var (
	Separator       = ":"
	SystemParameter = "system-parameter"
)

func getAppName() string {
	configs.BuildConfigEnv()
	return viper.GetString("application.name")
}

func CacheKeySysParamWithId(id int) string {
	return strings.Join([]string{getAppName(), SystemParameter, strconv.Itoa(id)}, Separator)
}

func CacheKeySysParams() string {
	return strings.Join([]string{getAppName(), SystemParameter, "list"}, Separator)
}
