package controller

import (
	"github.com/labstack/echo/v4"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//handler2 "myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/service"
)

func (controller *HelloWorldsController) AddRoutes(e *echo.Echo) {
	e.GET("/hello", controller.Hello)
}
