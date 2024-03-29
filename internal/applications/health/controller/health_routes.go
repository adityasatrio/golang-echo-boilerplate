package controller

import (
	"github.com/labstack/echo/v4"
)

func (controller *HealthController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName)
	group.GET("/health", controller.Health)
	group.GET("/health/database", controller.HealthDatabase)
	group.GET("/health/cache", controller.HealthCache)
}
