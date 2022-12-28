package main

import (
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	config "myapp/config"
	"myapp/internal/adapter/http"
	"myapp/internal/commons/middlewares"
	"myapp/internal/initialization"
	//"net/http"
)

/*type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	//set the client to the variable defined in package config
	//this will enable the client intance to be accessed anywhere through the accessor which is a function
	//named GetClient
	config.SetClient(client)

	//TODO : need to fix validator nya tidak berjalan
	initialization.SetupValidator(e)

	//add middlewares
	middlewares.InitMiddlewares(e)

	//http_routes
	//usecaseSysParam := usecase.NewUseCase()
	//handler.NewHandler(usecaseSysParam).AddRoutes(e)

	//define handler moved on this file
	http.SetupRouteHandler(e)

	//load config
	port := viper.GetString("application.port")
	err = e.Start(":" + port)
	if err != nil {
		return
	}

	e.Logger.Fatal(err)

}
