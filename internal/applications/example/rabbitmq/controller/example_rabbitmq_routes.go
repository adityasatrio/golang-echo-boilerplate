package controller

import "github.com/labstack/echo/v4"

func (ctl *ExampleRabbitMQController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName)
	group.GET("/example/rabbitmq", ctl.PublishRabbitMQ)
}
