package controller

import (
	"github.com/labstack/echo/v4"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//handler2 "myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/service"
)

func (c *SystemParameterController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName + "/system-parameter")

	group.POST("", c.Create)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
	group.GET("/:id", c.GetById)
	group.GET("", c.GetAll)

}
