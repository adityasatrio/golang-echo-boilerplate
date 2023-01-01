package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	config "myapp/config"
	database "myapp/config/database"
	"myapp/config/middleware"
	"myapp/config/validator"
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
	//from docs define close on this function, but will impact cant create DB session on repository
	defer dbConnection.Close()

	restApi.SetupRouteHandler(e, dbConnection)

	//load config
	port := viper.GetString("application.port")
	err := e.Start(":" + port)
	if err != nil {
		return
	}

	e.Logger.Fatal(err)
}
