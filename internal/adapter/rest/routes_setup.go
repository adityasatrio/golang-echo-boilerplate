package rest

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
	"myapp/configs/rabbitmq/connection"
	"myapp/ent"
	exampleRabbit "myapp/internal/applications/example/rabbitmq/controller"
	"myapp/internal/applications/health"
	"myapp/internal/applications/health/controller"
	quotes "myapp/internal/applications/quotes"
	quotesController "myapp/internal/applications/quotes/controller"
	"myapp/internal/applications/system_parameter"
	systemParameterController "myapp/internal/applications/system_parameter/controller"
	"myapp/internal/applications/user"
	userController "myapp/internal/applications/user/controller"
	"myapp/internal/component/rabbitmq/producer"
)

func SetupRouteHandler(e *echo.Echo, connDb *ent.Client, redisClient *redis.Client, connRabbitMQ *connection.RabbitMQConnection) {

	appName := viper.GetString("application.name")

	// Swagger OpenAPI Docs
	e.GET(appName+"/swagger/*", echoSwagger.WrapHandler)

	//manual injection
	helloWorldsService := health.InitializeHealthService(connDb, redisClient)
	controller.
		NewHealthController(helloWorldsService).
		AddRoutes(e, appName)

	//injection using code gen - google wire
	systemParameterService := system_parameter.InitializedSystemParameterService(connDb, redisClient)
	systemParameterController.
		NewSystemParameterController(systemParameterService).
		AddRoutes(e, appName)

	userService := user.InitializedUserService(connDb, redisClient)
	userController.
		NewUserController(userService).
		AddRoutes(e, appName)

	quotesService := quotes.InitializedQuotesService()
	quotesController.
		NewQuotesController(quotesService).
		AddRoutes(e, appName)

	//add example controller:
	producerService := producer.InitializedProducer(connRabbitMQ)
	exampleRabbit.NewExampleRabbitMQController(producerService).AddRoutes(e, appName)
}
