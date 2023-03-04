package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/helper/response"
	"myapp/internal/applications/post/dto"
	"myapp/internal/applications/post/service"
	"strconv"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{service: service}
}

func (c *PostController) Create(ctx echo.Context) error {
	request := new(dto.PostRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}

	err = ctx.Validate(request)
	if err != nil {
		return err
	}

	created, err := c.service.Create(ctx.Request().Context(), *request)
	if err != nil {
		return err
	}

	return response.Success(ctx, created)
}

func (c *PostController) Update(ctx echo.Context) error {
	request := new(dto.PostRequest)
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

	updated, err := c.service.Update(ctx.Request().Context(), *request, id)
	if err != nil {
		return err
	}

	return response.Success(ctx, updated)
}

func (c *PostController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	deleted, err := c.service.SoftDelete(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return response.Success(ctx, deleted)
}

func (c *PostController) GetById(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	result, err := c.service.GetById(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return response.Success(ctx, result)
}

func (c *PostController) GetAll(ctx echo.Context) error {

	result, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	return response.Success(ctx, result)
}
