package applications

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics
}
