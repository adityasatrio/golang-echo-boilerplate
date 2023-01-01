package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helper/response"
	"myapp/internal/applications/hello_worlds/service"
)

type HelloWorldsController struct {
	service service.HelloWorldsService
}

func NewHelloWorldsController(service service.HelloWorldsService) *HelloWorldsController {
	return &HelloWorldsController{
		service: service,
	}
}

func (controller *HelloWorldsController) Hello(c echo.Context) error {

	errorFlag := c.QueryParam("error")
	messageController := "hello from controller -"

	result, err := controller.service.Hello(c.Request().Context(), messageController, errorFlag)
	if err != nil {
		return err
	}

	return response.Success(c, result)
}
