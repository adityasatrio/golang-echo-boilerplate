package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helper"
	"myapp/helper/response"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/service"
	"strconv"
)

type SystemParameterController struct {
	service service.SystemParameterService
}

func NewSystemParameterController(service service.SystemParameterService) *SystemParameterController {
	return &SystemParameterController{
		service: service,
	}
}

func (c *SystemParameterController) Create(ctx echo.Context) error {
	request := new(dto.SystemParameterCreateRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}

	err = ctx.Validate(request)
	if err != nil {
		return err
	}

	created, err := c.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.SystemParameterResponse)
	err = helper.Mapper(&responseDto, created)
	if err != nil {
		return err
	}

	return response.Created(ctx, responseDto)
}

func (c *SystemParameterController) Update(ctx echo.Context) error {
	request := new(dto.SystemParameterUpdateRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}

	err = ctx.Validate(request)
	if err != nil {
		return err
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	updated, err := c.service.Update(ctx.Request().Context(), id, request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.SystemParameterResponse)
	err = helper.Mapper(&responseDto, updated)
	if err != nil {
		return err
	}

	return response.Success(ctx, updated)
}

func (c *SystemParameterController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	deleted, err := c.service.Delete(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	var responseDto = new(dto.SystemParameterResponse)
	err = helper.Mapper(&responseDto, deleted)
	if err != nil {
		return err
	}

	return response.Success(ctx, deleted)
}

func (c *SystemParameterController) GetById(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	result, err := c.service.GetById(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	var responseDto = new(dto.SystemParameterResponse)
	err = helper.Mapper(&responseDto, result)
	if err != nil {
		return err
	}

	return response.Success(ctx, result)
}

func (c *SystemParameterController) GetAll(ctx echo.Context) error {
	results, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	var responseDtos []*dto.SystemParameterResponse
	for _, result := range results {
		responseDto := new(dto.SystemParameterResponse)
		err = helper.Mapper(responseDto, result)
		if err != nil {
			return err
		}
		responseDtos = append(responseDtos, responseDto)
	}

	return response.Success(ctx, results)
}
