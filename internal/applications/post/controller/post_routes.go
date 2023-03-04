package controller

import "github.com/labstack/echo/v4"

func (c *PostController) AddRoutes(e *echo.Echo) {
	e.POST("/post", c.Create)
	e.PUT("/post/:id", c.Update)
	e.DELETE("/post/:id", c.Delete)
	e.GET("/post", c.GetAll)
	e.GET("/post/:id", c.GetById)
}
