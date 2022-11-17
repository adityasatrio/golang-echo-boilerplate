package http

import (
	"myapp/internal/applications/system_parameter/handler"
	"myapp/internal/applications/system_parameter/usecase"
)

func InitSystemParameterRoutes(e *echo.Echo) {

	//need refactor, should be more elegant
	h := handler.NewHandler(usecase.NewUseCase())
	e.GET("/hello", h.Hello)

	e.POST("/system-parameter", h.Create)
	e.PUT("/system-parameter/:id", h.Update)
	e.GET("/system-parameter/:id", h.GetById)
	e.DELETE("/system-parameter/:id", h.Delete)

}
