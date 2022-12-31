package main

import (
	"fmt"
	"myapp/config/middleware"
	"myapp/config/validator"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	config "myapp/config/database"
	restApi "myapp/internal/adapter/rest_api"
	//"net/rest_api"
)

/*type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(rest_api.StatusBadRequest, err.Error())
	}
	return nil
}*/

func main() {
	e := echo.New()

	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	//TODO : need to fix validator nya tidak berjalan
	validator.SetupValidator(e)
	validator.SetupHttpErrorHandler(e)

	//add middlewares
	middleware.InitMiddlewares(e)

	dbConnection := config.NewSqlEntClient()
	//dbConnection := config.NewEntClient()
	fmt.Println("dbConnection", dbConnection)

	restApi.SetupRouteHandler(e, dbConnection)

	//load config
	port := viper.GetString("application.port")
	err = e.Start(":" + port)
	if err != nil {
		return
	}

	e.Logger.Fatal(err)
}
