package controller

import "github.com/labstack/echo/v4"

func (c *UserController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName + "/user")

	group.POST("", c.Create)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
	group.GET("/:id", c.GetById)
	group.GET("", c.GetAll)

}
