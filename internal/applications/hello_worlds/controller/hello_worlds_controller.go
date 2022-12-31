package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helper/response"
	"myapp/internal/applications/hello_worlds/service"
)

type HelloWorldController struct {
	service service.HelloWorldService
}

func NewHelloWorldController(service service.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{
		service: service,
	}
}

func (controller *HelloWorldController) Hello(c echo.Context) error {

	errorFlag := c.QueryParam("error")
	messageController := "hello from controller -"

	result, err := controller.service.Hello(c.Request().Context(), messageController, errorFlag)
	if err != nil {
		return err
	}

	return response.Success(c, result)
}