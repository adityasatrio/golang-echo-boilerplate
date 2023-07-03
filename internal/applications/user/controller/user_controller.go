package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/user/dto"
	"myapp/internal/applications/user/service"
	"myapp/internal/helper"
	"myapp/internal/helper/response"
	"strconv"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Create(ctx echo.Context) error {
	request := new(dto.UserRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		//return helper.BadRequest(ctx, err)
		return err
	}

	created, err := c.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = helper.Mapper(&responseDto, created)
	if err != nil {
		return err
	}

	return response.Created(ctx, responseDto)
}

func (c *UserController) Update(ctx echo.Context) error {
	request := new(dto.UserRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	updated, err := c.service.Update(ctx.Request().Context(), uint64(id), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = helper.Mapper(&responseDto, updated)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

func (c *UserController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	deleted, err := c.service.Delete(ctx.Request().Context(), uint64(id))
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = helper.Mapper(&responseDto, deleted)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

func (c *UserController) GetById(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	result, err := c.service.GetById(ctx.Request().Context(), uint64(id))
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = helper.Mapper(&responseDto, result)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

func (c *UserController) GetAll(ctx echo.Context) error {
	results, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	var responseDtos []*dto.UserResponse
	for _, result := range results {
		responseDto := new(dto.UserResponse)
		err = helper.Mapper(responseDto, result)
		if err != nil {
			return err
		}
		responseDtos = append(responseDtos, responseDto)
	}

	return response.Success(ctx, results)
}
