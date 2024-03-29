package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

const (
	fileConfigType = ".env"
)

func InitGeneralEnv(e *echo.Echo) {
	viper.SetConfigFile(fileConfigType)
	viper.AddConfigPath(".")

	//set default variable for undefined on .env
	setDefaultKeys()

	log.Debugf("credential file : " + viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
		panic(e)
	}

	viper.WatchConfig()
	log.Infof("initialized WatchConfig(): success", "/."+fileConfigType)
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed:", e.Name)
	})

	log.Infof("initialized configs viper: success", "/."+fileConfigType)

}

func setDefaultKeys() {
	viper.SetDefault("application.port", 8080)
	viper.SetDefault("application.health.url", "/health")
	viper.SetDefault("application.mode", "dev")

	host := []string{"localhost", "https://labstack.com", "https://labstack.net"}
	viper.SetDefault("application.cors.allowedHost", host)

	viper.SetDefault("cache.configs.redis.host", "127.0.0.1")
	viper.SetDefault("cache.configs.redis.port", 6379)

	viper.SetDefault("cache.ttl.short", "3h")
	viper.SetDefault("cache.ttl.medium", "24h")
	viper.SetDefault("cache.ttl.long", "3d")

	viper.SetDefault("db.configs.maxIdleConn", 5)
	viper.SetDefault("db.configs.maxOpenConn", 10)

	viper.SetDefault("swagger.host", "localhost:8888")
	viper.SetDefault("rabbitmq.configs.recovery", 30)

	log.Infof("initialized default configs value : success")
	for key, value := range viper.AllSettings() {
		log.Infof("initialized default configs value %s : %s\n", key, value)
	}
}
