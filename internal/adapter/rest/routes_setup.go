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
	"myapp/middleware"
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

	// Auth-aware middleware for protected API routes
	authSvc := container.BuildAuthService()
	requireAuth := middleware.RequireAuth(authSvc)
	requireAdmin := middleware.RequireAdmin()

	// System parameter service: readable by Admin + User, writes Admin-only (enforced per-route)
	systemParameterService := container.BuildSystemParameterService()
	systemParameterController.
		NewSystemParameterController(systemParameterService).
		AddRoutes(e, appName, requireAuth)

	// User service: Admin-only, User role has zero access
	userService := container.BuildUserService()
	userController.
		NewUserController(userService).
		AddRoutes(e, appName, requireAuth, requireAdmin)

	// Quotes service
	quotesService := container.BuildQuotesService()
	quotesController.
		NewQuotesController(quotesService).
		AddRoutes(e, appName)

	// RabbitMQ producer
	producerService := container.BuildProducer()
	exampleRabbit.NewExampleRabbitMQController(producerService).AddRoutes(e, appName)
}
