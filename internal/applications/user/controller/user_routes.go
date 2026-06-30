package controller

import "github.com/labstack/echo/v4"

// AddRoutes registers the user management API. mw is applied to the whole
// group, typically auth + admin-only guards, since Admin has full access and
// the User role has zero access to user management.
func (c *UserController) AddRoutes(e *echo.Echo, appName string, mw ...echo.MiddlewareFunc) {
	group := e.Group(appName+"/user", mw...)

	group.POST("", c.Create)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
	group.GET("/:id", c.GetById)
	group.GET("", c.GetAll)

}
