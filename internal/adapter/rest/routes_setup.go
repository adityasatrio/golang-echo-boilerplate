package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
	exampleRabbit "myapp/internal/applications/example/rabbitmq/controller"
	healthController "myapp/internal/applications/health/controller"
	quotesController "myapp/internal/applications/quotes/controller"
	systemParameterController "myapp/internal/applications/system_parameter/controller"
	userController "myapp/internal/applications/user/controller"
	"myapp/internal/builder"
)

func SetupRouteHandler(e *echo.Echo, container *builder.Container) {

	appName := viper.GetString("application.name")

	// Swagger OpenAPI Docs
	e.GET(appName+"/swagger/*", echoSwagger.WrapHandler)

	// Health service
	healthService := container.BuildHealthService()
	healthController.
		NewHealthController(healthService).
		AddRoutes(e, appName)

	// System parameter service
	systemParameterService := container.BuildSystemParameterService()
	systemParameterController.
		NewSystemParameterController(systemParameterService).
		AddRoutes(e, appName)

	// User service
	userService := container.BuildUserService()
	userController.
		NewUserController(userService).
		AddRoutes(e, appName)

	// Quotes service
	quotesService := container.BuildQuotesService()
	quotesController.
		NewQuotesController(quotesService).
		AddRoutes(e, appName)

	// RabbitMQ producer
	producerService := container.BuildProducer()
	exampleRabbit.NewExampleRabbitMQController(producerService).AddRoutes(e, appName)
}
