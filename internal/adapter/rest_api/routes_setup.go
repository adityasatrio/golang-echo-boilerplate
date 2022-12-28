package rest_api

import (
	"github.com/labstack/echo/v4"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/handler"
	"myapp/internal/applications/system_parameter/repository"
	"myapp/internal/applications/system_parameter/service"
)

func SetupRouteHandler(e *echo.Echo, connection *ent.Client) {

	//systemParameterRepository := repository.NewSystemParameterRepository()
	//systemParameterUseCase := service.NewSystemParameterService(systemParameterRepository)
	systemParameterRepository := repository.NewSystemParameterRepository(connection)
	systemParameterUseCase := service.NewSystemParameterService(systemParameterRepository)
	handler.NewHandler(systemParameterUseCase).AddRoutes(e)

	//list down other route and handlers
}
