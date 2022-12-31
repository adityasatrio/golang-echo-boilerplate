package controller

import (
	"github.com/labstack/echo/v4"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//handler2 "myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/service"
)

func (c *SystemParameterController) AddRoutes(e *echo.Echo) {

	//need refactor, should be more elegant
	//c := handler2.NewSystemParameterController(service.NewUseCase())
	//e.GET("/hello", handler.Hello)

	e.POST("/system-parameter", c.Create)
	e.PUT("/system-parameter/:id", c.Update)
	e.DELETE("/system-parameter/:id", c.Delete)
	e.GET("/system-parameter/:id", c.GetById)
	e.GET("/system-parameter", c.GetAll)

}
