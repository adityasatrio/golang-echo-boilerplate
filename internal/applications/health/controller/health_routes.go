package controller

import (
	"github.com/labstack/echo/v4"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//handler2 "myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/service"
)

func (controller *HealthController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName)
	group.GET("/health", controller.Health)
}
