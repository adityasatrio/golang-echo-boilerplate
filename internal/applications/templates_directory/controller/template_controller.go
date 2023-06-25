package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/templates_directory/dto"
	"myapp/internal/applications/templates_directory/service"
	"myapp/internal/apputils"
	"myapp/internal/apputils/response"
)

type TemplateController struct {
	service service.TemplateService
}

func NewTemplateController(service service.TemplateService) *TemplateController {
	return &TemplateController{
		service: service,
	}
}

func (c *TemplateController) validateAndParseFunction(ctx echo.Context) error {
	//controller layer is only to handle request and response, you can do validate your request using validators
	// validation logic if needed. You are forbidden to put any business logic in this layer

	request := new(dto.ExampleRequest)
	err := apputils.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	//call service
	serviceResult := "service_result"

	//handle response after get result from service layer
	var responseDto = new(dto.ExampleResponse)
	err = apputils.Mapper(&responseDto, serviceResult)
	if err != nil {
		return err
	}

	return response.Created(ctx, responseDto)
}
