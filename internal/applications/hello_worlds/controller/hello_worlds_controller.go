package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/hello_worlds/dto"
	"myapp/internal/applications/hello_worlds/service"
	"myapp/internal/apputils"
	"time"
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
	messageController := "hello from controller ->"

	result, err := controller.service.Hello(c.Request().Context(), messageController, errorFlag)
	if err != nil {
		return err
	}

	var responseDto = dto.HelloWorldsResponse{
		Message:   result,
		CreatedBy: "user",
		CreatedAt: time.Now().String(),
		UpdatedBy: "user",
		UpdatedAt: time.Now().String(),
	}

	return apputils.Success(c, responseDto)
}
