package controller

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/exceptions"
	"myapp/helper/response"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/service"
	"strconv"

	//"myapp/pkg/validator"
	"net/http"
)

type SystemParameterController struct {
	response response.Response
	service  service.SystemParameterService
}

func NewSystemParameterController(response response.Response, service service.SystemParameterService) *SystemParameterController {
	return &SystemParameterController{
		response: response,
		service:  service,
	}
}

func (c *SystemParameterController) Create(ctx echo.Context) error {
	request := new(dto.SystemParameterCreateRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, err)
	}

	//err = validator.ReqBody(c, request)
	err = ctx.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, err)
	}

	created, err := c.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return c.response.Error(ctx, http.StatusInternalServerError, err)
	}

	return c.response.Created(ctx, created)
}

func (c *SystemParameterController) Update(ctx echo.Context) error {
	request := new(dto.SystemParameterUpdateRequest)
	err := ctx.Bind(&request)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, err)
	}

	//err = validator.ReqBody(c, request)
	err = ctx.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, err)
	}

	//TODO : create helper for get param and validate
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, err)
	}

	created, err := c.service.Update(ctx.Request().Context(), id, request)
	if err != nil {
		return c.response.Error(ctx, http.StatusInternalServerError, err)
	}

	return c.response.Success(ctx, created)
}

/*func (handler *SystemParameterController) Update(c echo.Context) error {
	id := c.Param("id")
	var request dto.SystemParameterRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(id)
	return c.JSON(http.StatusCreated, request)

}*/

func (c *SystemParameterController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.response.Error(ctx, http.StatusBadRequest, err)
	}

	created, err := c.service.Delete(ctx.Request().Context(), id)
	if errors.Is(err, exceptions.TargetDataGetError) {
		return c.response.Base(ctx, http.StatusNotFound, response.StatusFailed, created, err)

	} else if errors.Is(err, exceptions.TargetDataDeleteError) {
		return c.response.Error(ctx, http.StatusInternalServerError, err)
	}

	return c.response.Success(ctx, created)
}

func (c *SystemParameterController) GetById(ctx echo.Context) error {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.response.Base(ctx, http.StatusBadRequest, response.StatusFailed, nil, err)
	}

	created, err := c.service.GetById(ctx.Request().Context(), id)
	if err != nil {
		return c.response.Base(ctx, http.StatusInternalServerError, response.StatusFailed, created, err)
	}

	return c.response.Base(ctx, http.StatusOK, response.StatusSuccess, created, err)
}

func (c *SystemParameterController) GetAll(ctx echo.Context) error {
	created, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return c.response.Base(ctx, http.StatusInternalServerError, response.StatusFailed, created, err)
	}

	return c.response.Base(ctx, http.StatusOK, response.StatusSuccess, created, err)
}
