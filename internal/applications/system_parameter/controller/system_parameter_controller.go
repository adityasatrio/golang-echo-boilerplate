package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/service"
	"myapp/internal/helper"
	"myapp/internal/helper/response"
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

// Create is controller to create new system parameter.
//
//	@summary		Create new system parameter
//	@description	Create new system parameter
//	@tags			system parameter
//	@accept			json
//	@produce		json
//	@param			body	body		dto.SystemParameterCreateRequest	true	"Create System Parameter DTO"
//	@success		201		{object}	response.body{data=dto.SystemParameterResponse}
//	@failure		422		{object}	response.body
//	@failure		500		{object}	response.body
//	@router			/system-parameter [post]
func (c *SystemParameterController) Create(ctx echo.Context) error {
	request := new(dto.SystemParameterCreateRequest)
	err := helper.BindAndValidate(ctx, request)
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

// Update is controller to update existing system parameter.
//
//	@summary		Update a system parameter
//	@description	Update a system parameter
//	@tags			system parameter
//	@accept			json
//	@produce		json
//	@param			id		path		int									true	"Existing system parameter ID"
//	@param			body	body		dto.SystemParameterUpdateRequest	true	"Update System Parameter DTO"
//	@success		200		{object}	response.body{data=dto.SystemParameterResponse}
//	@failure		404		{object}	response.body
//	@failure		422		{object}	response.body
//	@failure		500		{object}	response.body
//	@router			/system-parameter/{id} [put]
func (c *SystemParameterController) Update(ctx echo.Context) error {
	request := new(dto.SystemParameterUpdateRequest)
	err := helper.BindAndValidate(ctx, request)

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

	return response.Success(ctx, responseDto)
}

// Delete is controller to delete existing system parameter.
//
//	@summary		Delete a system parameter
//	@description	Delete a system parameter
//	@tags			system parameter
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Existing system parameter ID"
//	@success		200	{object}	response.body{data=dto.SystemParameterResponse}
//	@failure		404	{object}	response.body
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/system-parameter/{id} [delete]
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

	return response.Success(ctx, responseDto)
}

// GetById is controller to get system parameter by its id.
//
//	@summary		Get a system parameter by id
//	@description	Get a system parameter by id
//	@tags			system parameter
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Existing system parameter ID"
//	@success		200	{object}	response.body{data=dto.SystemParameterResponse}
//	@failure		404	{object}	response.body
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/system-parameter/{id} [get]
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

	return response.Success(ctx, responseDto)
}

// GetAll is controller to get all system parameters.
//
//	@summary		Get all system parameters
//	@description	Get all system parameters
//	@tags			system parameter
//	@accept			json
//	@produce		json
//	@success		200	{object}	response.body{data=[]dto.SystemParameterResponse}
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/system-parameter [get]
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

	return response.Success(ctx, responseDtos)
}
