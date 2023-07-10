package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/quotes/dto"
	"myapp/internal/applications/quotes/service"
	"myapp/internal/helper"
	"myapp/internal/helper/response"
)

type QuotesController struct {
	service service.QuotesService
}

func NewQuotesController(service service.QuotesService) *QuotesController {
	return &QuotesController{
		service: service,
	}
}

func (c *QuotesController) Get(ctx echo.Context) error {
	query1 := ctx.QueryParam("query1")
	query2 := ctx.QueryParam("query2")

	queryParameter := map[string]string{"query1": query1, "query2": query2}
	result, err := c.service.GetQuotes(ctx.Request().Context(), queryParameter)
	if err != nil {
		return err
	}

	var responseDto = new(dto.QuoteResponse)
	responseDto.Custom = "beyond limit"
	err = helper.Mapper(&responseDto, result)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

func (c *QuotesController) Post(ctx echo.Context) error {
	request := new(dto.QuoteRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	result, err := c.service.PostQuotes(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.QuoteResponse)
	responseDto.Custom = "beyond limit"
	err = helper.Mapper(&responseDto, result)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}
