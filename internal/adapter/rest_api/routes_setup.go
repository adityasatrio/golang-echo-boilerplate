package rest_api

import (
	"github.com/labstack/echo/v4"
	"myapp/ent"
	helloController "myapp/internal/applications/hello_worlds/controller"
	"myapp/internal/applications/hello_worlds/repository"
	"myapp/internal/applications/hello_worlds/service"
	"myapp/internal/applications/system_parameter"
	systemParameterController "myapp/internal/applications/system_parameter/controller"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client) {

	//manual injection
	helloWorldsRepository := repository.NewHelloWorldsRepository(connDb)
	helloWorldsService := service.NewHelloWorldsService(helloWorldsRepository)
	helloController.
		NewHelloWorldsController(helloWorldsService).
		AddRoutes(e)

	//injection using code gen - google wire
	SystemParameterService := system_parameter.InitializedSystemParameterService(connDb)
	systemParameterController.NewSystemParameterController(SystemParameterService).
		AddRoutes(e)

}
