package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/middleware"
)

// AddRoutes registers the system parameter API. mw (typically auth-only) is
// applied to the whole group since both Admin and User can read; write
// operations additionally require the Admin role.
func (c *SystemParameterController) AddRoutes(e *echo.Echo, appName string, mw ...echo.MiddlewareFunc) {
	group := e.Group(appName+"/system-parameter", mw...)

	group.POST("", c.Create, middleware.RequireAdmin())
	group.PUT("/:id", c.Update, middleware.RequireAdmin())
	group.DELETE("/:id", c.Delete, middleware.RequireAdmin())
	group.GET("/:id", c.GetById)
	group.GET("", c.GetAll)

}
