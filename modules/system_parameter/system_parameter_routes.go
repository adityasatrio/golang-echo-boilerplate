package system_parameter

import (
	"github.com/labstack/echo/v4"
	"myapp/modules/system_parameter/handler"
	"myapp/modules/system_parameter/usecase"
)

func InitSystemParameterRoutes(e *echo.Echo) {

	h := handler.New(usecase.New())
	e.GET("/system-parameter", h.CreateSystemParameter)

	/*e.POST("/system-parameter", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.PUT("/system-parameter/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/system-parameter/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.DELETE("/system-parameter/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})*/

}
