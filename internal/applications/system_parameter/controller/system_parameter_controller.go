package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/helper/response"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/service"
	"strconv"

	//"myapp/pkg/validator"
	"net/http"
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
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	//err = validator.ReqBody(c, request)
	err = ctx.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	created, err := c.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return response.ServiceErrorHandler(ctx, created, err)
	}

	return response.Created(ctx, created)
}

func (c *SystemParameterController) Update(ctx echo.Context) error {
	request := new(dto.SystemParameterUpdateRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	//err = validator.ReqBody(c, request)
	err = ctx.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	//TODO : create helper for get param and validate
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	updated, err := c.service.Update(ctx.Request().Context(), id, request)
	if err != nil {
		return response.ServiceErrorHandler(ctx, updated, err)
	}

	return response.Success(ctx, updated)
}

func (c *SystemParameterController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	deleted, err := c.service.Delete(ctx.Request().Context(), id)
	if err != nil {
		return response.ServiceErrorHandler(ctx, deleted, err)
	}

	return response.Success(ctx, deleted)
}

func (c *SystemParameterController) GetById(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	result, err := c.service.GetById(ctx.Request().Context(), id)
	if err != nil {
		return response.ServiceErrorHandler(ctx, result, err)
	}

	return response.Success(ctx, result)
}

func (c *SystemParameterController) GetAll(ctx echo.Context) error {
	results, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return response.ServiceErrorHandler(ctx, results, err)
	}

	return response.Success(ctx, results)
}
