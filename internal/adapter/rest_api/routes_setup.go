package rest_api

import (
	"github.com/labstack/echo/v4"
	"myapp/ent"
	"myapp/internal/applications"
	helloController "myapp/internal/applications/hello_worlds/controller"
	"myapp/internal/applications/hello_worlds/repository"
	"myapp/internal/applications/hello_worlds/service"
	"myapp/internal/applications/system_parameter"
	systemParameterController "myapp/internal/applications/system_parameter/controller"
	"myapp/internal/applications/user"
	userController "myapp/internal/applications/user/controller"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client) {
	//apps status metrics
	applications.AddRoutes(e)

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

	UserService := user.InitializedUserService(connDb)
	userController.NewUserController(UserService).AddRoutes(e)

}
