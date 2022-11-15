package system_parameter

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitSystemParameterRoutes(e *echo.Echo) {

	e.POST("/system-parameter", func(c echo.Context) error {
		c.Request().Context()
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
	})
}
