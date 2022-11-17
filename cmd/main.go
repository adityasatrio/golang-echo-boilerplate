package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	apihttp "myapp/internal/adapter/api_http/system_parameter"
	"myapp/internal/commons/middlewares"
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
	//e.HTTPErrorHandler = global_handler.InitHttpErrorHandler()

	e.HTTPErrorHandler = func(err error, context echo.Context) {

	}

	//add middlewares
	middlewares.InitMiddlewares(e)

	//http_routes
	apihttp.InitSystemParameterRoutes(e)

	//load config
	err := e.Start(":1234")
	if err != nil {
		return
	}

	e.Logger.Fatal(err)

}
