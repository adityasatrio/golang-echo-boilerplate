package controller

import "github.com/labstack/echo/v4"

func (c *PostController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName + "/post")
	//All routes belong here, this is only post, please modify based on your needs

	group.POST("", c.validateAndParseFunction)
	group.PUT("/:id", c.validateAndParseFunction)
	group.DELETE("/:id", c.validateAndParseFunction)
	group.GET("/:id", c.validateAndParseFunction)
	group.GET("", c.validateAndParseFunction)
}
