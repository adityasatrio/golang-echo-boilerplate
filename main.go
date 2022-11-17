package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"myapp/commons/middlewares"
	"myapp/modules/system_parameter"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

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
