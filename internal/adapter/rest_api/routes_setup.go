package rest_api

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"myapp/ent"
	helloController "myapp/internal/applications/hello_worlds/controller"
	"myapp/internal/applications/hello_worlds/repository"
	"myapp/internal/applications/hello_worlds/service"
	"myapp/internal/applications/system_parameter"
	systemParameterController "myapp/internal/applications/system_parameter/controller"
	"myapp/internal/applications/user"
	userController "myapp/internal/applications/user/controller"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client, redisClient *redis.Client) {

	appName := viper.GetString("application.name")

	//manual injection
	helloWorldsRepository := repository.NewHelloWorldsRepository(connDb)
	helloWorldsService := service.NewHelloWorldsService(helloWorldsRepository)
	helloController.
		NewHelloWorldsController(helloWorldsService).
		AddRoutes(e, appName)

	//injection using code gen - google wire
	SystemParameterService := system_parameter.InitializedSystemParameterService(connDb, redisClient)
	systemParameterController.
		NewSystemParameterController(SystemParameterService).
		AddRoutes(e, appName)

	UserService := user.InitializedUserService(connDb, redisClient)
	userController.
		NewUserController(UserService).
		AddRoutes(e, appName)

}
