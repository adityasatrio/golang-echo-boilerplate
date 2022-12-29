package rest_api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/commons/response"
	"myapp/ent"

	helloController "myapp/internal/applications/hello_worlds/controller"
	helloRepository "myapp/internal/applications/hello_worlds/repository"
	helloService "myapp/internal/applications/hello_worlds/service"

	"myapp/internal/applications/system_parameter/controller"
	"myapp/internal/applications/system_parameter/repository"
	"myapp/internal/applications/system_parameter/service"
)

func SetupRouteHandler(e *echo.Echo, connection *ent.Client) {

	response := response.NewBaseResponse()

	helloWorldsRepository := helloRepository.NewHelloWorldsRepository(connection)
	helloWorldsService := helloService.NewHelloWorldsService(helloWorldsRepository)

	helloController.
		NewHelloWorldController(response, helloWorldsService).
		AddRoutes(e)

	systemParameterRepository := repository.NewSystemParameterRepository(connection)
	fmt.Println("systemParameterRepository", systemParameterRepository)

	systemParameterUseCase := service.NewSystemParameterService(systemParameterRepository)
	fmt.Println("systemParameterUseCase", systemParameterUseCase)

	controller.NewHandler(systemParameterUseCase).AddRoutes(e)

}
