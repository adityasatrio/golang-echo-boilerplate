package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"myapp/ent"
	helloController "myapp/internal/applications/health/controller"
	"myapp/internal/applications/health/repository"
	"myapp/internal/applications/health/service"
	"myapp/internal/applications/system_parameter"
	systemParameterController "myapp/internal/applications/system_parameter/controller"
	"myapp/internal/applications/user"
	userController "myapp/internal/applications/user/controller"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client) {

	appName := viper.GetString("application.name")

	//manual injection
	helloWorldsRepository := repository.NewHealthRepository(connDb)
	helloWorldsService := service.NewHealthService(helloWorldsRepository)
	helloController.
		NewHealthController(helloWorldsService).
		AddRoutes(e, appName)

	//injection using code gen - google wire
	SystemParameterService := system_parameter.InitializedSystemParameterService(connDb)
	systemParameterController.
		NewSystemParameterController(SystemParameterService).
		AddRoutes(e, appName)

	UserService := user.InitializedUserService(connDb)
	userController.
		NewUserController(UserService).
		AddRoutes(e, appName)

}
