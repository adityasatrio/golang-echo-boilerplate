package controller

import (
	"github.com/labstack/echo/v4"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/controller"
	//handler2 "myapp/internal/applications/system_parameter/controller"
	//"myapp/internal/applications/system_parameter/service"
)

func (controller *SystemParameterController) AddRoutes(e *echo.Echo) {

	//need refactor, should be more elegant
	//controller := handler2.NewSystemParameterController(service.NewUseCase())
	//e.GET("/hello", handler.Hello)

	//e.POST("/system-parameter", handler.Create)
	//e.PUT("/system-parameter/:id", handler.Update)
	e.GET("/system-parameter/:id", controller.GetById)
	e.DELETE("/system-parameter/:id", controller.Delete)
}
