package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helper/response"
	"myapp/internal/applications/user/dto"
	"myapp/internal/applications/user/service"
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
	err := ctx.Bind(request)
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

	return response.Created(ctx, created)
}

func (c *UserController) Update(ctx echo.Context) error {
	request := new(dto.UserRequest)
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

	updated, err := c.service.Update(ctx.Request().Context(), uint64(id), request)
	if err != nil {
		return err
	}

	return response.Success(ctx, updated)
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

	return response.Success(ctx, deleted)
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

	return response.Success(ctx, result)
}

func (c *UserController) GetAll(ctx echo.Context) error {
	results, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	return response.Success(ctx, results)
}
