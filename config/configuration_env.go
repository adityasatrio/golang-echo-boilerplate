package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
)

const (
	fileConfigPath = "."
	fileConfigType = "yml"
	fileConfigName = "app.config"
)

func SetupConfigEnv(e *echo.Echo) {
	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigType(fileConfigType)
	viper.SetConfigName(fileConfigName)

	setDefaultKeys()
	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
		panic(e)
	}

	log.Default().Println("initialized config viper: success", fileConfigPath+"/"+fileConfigName+"."+fileConfigType)

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	log.Default().Println("initialized WatchConfig(): success", fileConfigPath+"/"+fileConfigName+"."+fileConfigType)
}

func setDefaultKeys() {
	viper.SetDefault("application.port", 8080)

	host := []string{"localhost", "https://labstack.com", "https://labstack.net"}
	viper.SetDefault("application.cors.allowedHost", host)

	viper.SetDefault("db.config.username", "root")
	///viper.SetDefault("db.config.password", "password")
	viper.SetDefault("db.config.host", "127.0.0.1")
	viper.SetDefault("db.config.port", "3306")
	viper.SetDefault("db.config.database", "echo_sample")

	viper.SetDefault("cache.config.ristretto.numCounters", 1000)
	viper.SetDefault("cache.config.ristretto.maxCost", 100)

	viper.SetDefault("cache.config.redis.username", "root")
	//viper.SetDefault("cache.config.redis.password", "password")
	viper.SetDefault("cache.config.redis.DB", 0)
	viper.SetDefault("cache.config.redis.poolSize", 10)

	viper.SetDefault("cache.config.redis.host", "127.0.0.1")
	viper.SetDefault("cache.config.redis.port", 6379)

}
