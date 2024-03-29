package controller

import "github.com/labstack/echo/v4"

func (c *TemplateController) AddRoutes(e *echo.Echo) {
	//All routes belong here, this is only template, please modify based on your needs

	e.POST("/template", c.validateAndParseFunction)
	e.PUT("/template/:id", c.validateAndParseFunction)
	e.DELETE("/template/:id", c.validateAndParseFunction)
	e.GET("/template/:id", c.validateAndParseFunction)
	e.GET("/template", c.validateAndParseFunction)
}
