package controller

import "github.com/labstack/echo/v4"

func (c *UserController) AddRoutes(e *echo.Echo) {

	e.POST("/user", c.Create)
	e.PUT("/user/:id", c.Update)
	e.DELETE("/user/:id", c.Delete)
	e.GET("/user/:id", c.GetById)
	e.GET("/user", c.GetAll)

}
