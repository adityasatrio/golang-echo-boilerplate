package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/post/dto"
	"myapp/internal/applications/post/service"
	"myapp/internal/helper"
	"myapp/internal/helper/response"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{
		service: service,
	}
}

func (c *PostController) validateAndParseFunction(ctx echo.Context) error {
	//TLDR; controller layer is only to handle input and output from external domain from this application commonly to handle API request and response.
	//If needed validation logic for request, you can do validation using validators
	//You are forbidden to put any business logic in this layer

	request := new(dto.ExampleRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	//call service
	serviceResult := "service_result"

	//handle response after get result from service layer
	var responseDto = new(dto.ExampleResponse)
	err = helper.Mapper(&responseDto, serviceResult)
	if err != nil {
		return err
	}

	return response.Created(ctx, responseDto)
}
