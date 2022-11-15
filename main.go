package main

import (
	"github.com/labstack/echo/v4"
	"myapp/commons/middlewares"
	"myapp/modules/system_parameter"
)

func main() {
	e := echo.New()

	//add middlewares
	middlewares.InitMiddlewares(e)

	//http_routes
	system_parameter.InitSystemParameterRoutes(e)

	//load config
	err := e.Start(":1234")
	if err != nil {
		return
	}

	e.Logger.Fatal(err)

}
