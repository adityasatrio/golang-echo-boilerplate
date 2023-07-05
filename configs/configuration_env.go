package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

const (
	FileConfigName = "app.config"
	FileConfigPath = "."
	FileConfigType = "yml"
)

func SetupConfigEnv(e *echo.Echo) {
	viper.AddConfigPath(FileConfigPath)
	viper.SetConfigType(FileConfigType)
	viper.SetConfigName(FileConfigName)

	setDefaultKeys()
	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
		panic(e)
	}

	log.Infof("initialized configs viper: success", FileConfigPath+"/"+FileConfigName+"."+FileConfigType)

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	log.Infof("initialized WatchConfig(): success", FileConfigPath+"/"+FileConfigName+"."+FileConfigType)

	//set value for global variable, when value in app.config:
	InitGlobalVariable()
}

func setDefaultKeys() {
	viper.SetDefault("application.port", 8080)

	host := []string{"localhost", "https://labstack.com", "https://labstack.net"}
	viper.SetDefault("application.cors.allowedHost", host)

	//viper.SetDefault("db.configs.username", "root")
	//viper.SetDefault("db.configs.password", "password")
	viper.SetDefault("db.configs.host", "127.0.0.1")
	viper.SetDefault("db.configs.port", "3306")
	viper.SetDefault("db.configs.database", "echo_sample")

	viper.SetDefault("cache.configs.ristretto.numCounters", 1000)
	viper.SetDefault("cache.configs.ristretto.maxCost", 100)

	// viper.SetDefault("cache.configs.redis.username", "root")
	// viper.SetDefault("cache.configs.redis.password", "password")
	viper.SetDefault("cache.configs.redis.db", 0)
	viper.SetDefault("cache.configs.redis.poolSize", 10)

	viper.SetDefault("cache.configs.redis.host", "127.0.0.1")
	viper.SetDefault("cache.configs.redis.port", 6379)

	viper.SetDefault("cache.ttl.short-period", "3h")
	viper.SetDefault("cache.ttl.medium-period", "24h")
	viper.SetDefault("cache.ttl.long-period", "3d")

	log.Infof("initialized default configs value : success", viper.AllSettings())
}
