package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	config "myapp/config"
	"myapp/config/cache"
	database "myapp/config/database"
	"myapp/config/middleware"
	"myapp/config/validator"
	"myapp/ent"
	restApi "myapp/internal/adapter/rest_api"
)

func main() {
	e := echo.New()
	config.SetupConfigEnv(e)
	middleware.SetupMiddlewares(e)
	validator.SetupValidator(e)
	validator.SetupHttpErrorHandler(e)

	dbConnection := database.NewSqlEntClient() //using sqlDb wrapped by ent
	//dbConnection := config.NewEntClient() //using ent only
	log.Info("initialized database configuration=", dbConnection)
	defer func(dbConnection *ent.Client) {
		err := dbConnection.Close()
		if err != nil {
			log.Fatal("error when closed DB connection")
		}
	}(dbConnection)
	cacheManager := cache.NewCacheManager(e)

	restApi.SetupRouteHandler(e, dbConnection, cacheManager)

	//load config
	port := viper.GetString("application.port")
	err := e.Start(":" + port)
	if err != nil {
		return
	}

	e.Logger.Fatal(err)
}
