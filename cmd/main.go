package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"myapp/configs"
	"myapp/configs/database"
	"myapp/configs/validator"
	"myapp/ent"
	restApi "myapp/internal/adapter/rest_api"
	"myapp/middleware"
)

func main() {
	e := echo.New()

	configs.SetupConfigEnv(e)
	middleware.SetupMiddlewares(e)
	validator.SetupValidator(e)
	validator.SetupGlobalHttpUnhandleErrors(e)

	dbConnection := database.NewSqlEntClient() //using sqlDb wrapped by ent
	//dbConnection := database.NewEntClient() //using ent only

	log.Info("initialized database configuration=", dbConnection)
	//from docs define close on this function, but will impact cant create DB session on repository
	defer func(dbConnection *ent.Client) {
		err := dbConnection.Close()
		if err != nil {
			log.Fatalf("error initialized database configuration=", err)
		}
	}(dbConnection)

	restApi.SetupRouteHandler(e, dbConnection)

	//load configs
	port := viper.GetString("application.port")
	err := e.Start(":" + port)
	if err != nil {
		return
	}

	e.Logger.Fatal(err)
}
