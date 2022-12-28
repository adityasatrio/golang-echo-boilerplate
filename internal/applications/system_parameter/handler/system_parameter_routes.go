package handler

import (
	"github.com/labstack/echo/v4"
	//"myapp/internal/applications/system_parameter/handler"
	//"myapp/internal/applications/system_parameter/handler"
	//"myapp/internal/applications/system_parameter/handler"
	//handler2 "myapp/internal/applications/system_parameter/handler"
	//"myapp/internal/applications/system_parameter/service"
)

func (handler *SystemParameterHandler) AddRoutes(e *echo.Echo) {

	//need refactor, should be more elegant
	//handler := handler2.NewHandler(service.NewUseCase())
	e.GET("/hello", handler.Hello)

	e.POST("/system-parameter", handler.Create)
	e.PUT("/system-parameter/:id", handler.Update)
	e.GET("/system-parameter/:id", handler.GetById)
	e.DELETE("/system-parameter/:id", handler.Delete)
}
