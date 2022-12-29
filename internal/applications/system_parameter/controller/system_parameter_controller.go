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
	response response.Response
	service  service.SystemParameterService
}

func NewSystemParameterController(response response.Response, service service.SystemParameterService) *SystemParameterController {
	return &SystemParameterController{
		response: response,
		service:  service,
	}
}

func (controller *SystemParameterController) Create(c echo.Context) error {
	request := new(dto.SystemParameterCreateRequest)
	err := c.Bind(&request)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	//err = validator.ReqBody(c, request)
	err = c.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	created, err := controller.service.Create(c.Request().Context(), request)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusInternalServerError, response.FAILED, created, err)
	}

	return controller.response.BaseResponse(c, http.StatusOK, response.SUCCESS, created, err)
}

func (controller *SystemParameterController) Update(c echo.Context) error {
	request := new(dto.SystemParameterUpdateRequest)
	err := c.Bind(&request)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	//err = validator.ReqBody(c, request)
	err = c.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	created, err := controller.service.Update(c.Request().Context(), id, request)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusInternalServerError, response.FAILED, created, err)
	}

	return controller.response.BaseResponse(c, http.StatusOK, response.SUCCESS, created, err)
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

func (controller *SystemParameterController) Delete(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	created, err := controller.service.Delete(c.Request().Context(), id)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusInternalServerError, response.FAILED, created, err)
	}

	return controller.response.BaseResponse(c, http.StatusOK, response.SUCCESS, created, err)
}

func (controller *SystemParameterController) GetById(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusBadRequest, response.FAILED, nil, err)
	}

	created, err := controller.service.GetById(c.Request().Context(), id)
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusInternalServerError, response.FAILED, created, err)
	}

	return controller.response.BaseResponse(c, http.StatusOK, response.SUCCESS, created, err)
}

func (controller *SystemParameterController) GetAll(c echo.Context) error {
	created, err := controller.service.GetAll(c.Request().Context())
	if err != nil {
		return controller.response.BaseResponse(c, http.StatusInternalServerError, response.FAILED, created, err)
	}

	return controller.response.BaseResponse(c, http.StatusOK, response.SUCCESS, created, err)
}
