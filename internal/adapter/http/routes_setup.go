package http

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/system_parameter/handler"
	"myapp/internal/applications/system_parameter/usecase"
)

func SetupRouteHandler(e *echo.Echo) {

	//TODO : use case ini nanti di move ke struct penampung, seperti factory
	usecaseSysParam := usecase.NewUseCase()
	handler.NewHandler(usecaseSysParam).AddRoutes(e)

	//list down other route and handlers
}
