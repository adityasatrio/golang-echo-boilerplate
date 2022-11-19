package main

import (
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"myapp/internal/adapter/http"
	"myapp/internal/applications/system_parameter/handler"
	"myapp/internal/applications/system_parameter/usecase"
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
	initialization.SetupValidator(e)

	//add middlewares
	middlewares.InitMiddlewares(e)

	//http_routes
	usecaseSysParam := usecase.NewUseCase()
	handler.NewHandler(usecaseSysParam).AddRoutes(e)

	http.SetupRouteHandler(e)

	//load config
	err := e.Start(":1234")
	if err != nil {
		return
	}

	e.Logger.Fatal(err)

}
