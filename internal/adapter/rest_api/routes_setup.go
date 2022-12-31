package rest_api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/ent"
	helloController "myapp/internal/applications/hello_worlds/controller"
	helloRepository "myapp/internal/applications/hello_worlds/repository"
	helloService "myapp/internal/applications/hello_worlds/service"
	"myapp/internal/applications/system_parameter/repository/db"

	"myapp/internal/applications/system_parameter/controller"
	"myapp/internal/applications/system_parameter/service"
)

func SetupRouteHandler(e *echo.Echo, connection *ent.Client) {

	helloWorldsRepository := helloRepository.NewHelloWorldsRepository(connection)
	helloWorldsService := helloService.NewHelloWorldsService(helloWorldsRepository)

	helloController.
		NewHelloWorldController(helloWorldsService).
		AddRoutes(e)

	systemParameterRepository := db.NewSystemParameterRepository(connection)
	fmt.Println("systemParameterRepository", systemParameterRepository)

	systemParameterUseCase := service.NewSystemParameterService(systemParameterRepository)
	fmt.Println("systemParameterUseCase", systemParameterUseCase)

	controller.
		NewSystemParameterController(systemParameterUseCase).
		AddRoutes(e)

}
