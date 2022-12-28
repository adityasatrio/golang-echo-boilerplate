package http

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/system_parameter/handler"
	"myapp/internal/applications/system_parameter/repository"
	"myapp/internal/applications/system_parameter/usecase"
)

func SetupRouteHandler(e *echo.Echo) {

	//TODO : use case ini nanti di move ke struct penampung, seperti factory
	//TODO : possible to use wire as DI

	//systemParameterRepository := repository.NewRepository()
	//systemParameterUseCase := usecase.NewUseCase(systemParameterRepository)
	systemParameterRepository := repository.NewRepository()
	systemParameterUseCase := usecase.NewUseCase(systemParameterRepository)
	handler.NewHandler(systemParameterUseCase).AddRoutes(e)

	//list down other route and handlers
}
