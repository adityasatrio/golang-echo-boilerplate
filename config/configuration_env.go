package config

import (
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

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
		panic(e)
	}

	log.Default().Println("initialized config viper: success", fileConfigPath+"/"+fileConfigName+"."+fileConfigType)
}
