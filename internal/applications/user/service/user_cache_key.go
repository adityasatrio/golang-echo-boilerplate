package service

import (
	"github.com/spf13/viper"
	"myapp/configs"
	"strconv"
	"strings"
)

var (
	Separator = ":"
	User      = "user"
)

func getAppName() string {
	configs.BuildConfigEnv()
	return viper.GetString("application.name")
}

func CacheKeyUserWithId(id uint64) string {
	return strings.Join([]string{getAppName(), User, strconv.FormatUint(id, 10)}, Separator)
}

func CacheKeyUsers() string {
	return strings.Join([]string{getAppName(), User, "list"}, Separator)
}
