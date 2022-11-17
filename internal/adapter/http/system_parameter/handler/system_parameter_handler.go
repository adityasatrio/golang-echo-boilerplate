package handler

import (
	"fmt"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/usecase"
	"myapp/internal/commons/utils/validator"
	"net/http"
)

type SystemParameterHandler struct {
	useCase usecase.SystemParameterCase
}

func NewHandler(useCase usecase.SystemParameterCase) *SystemParameterHandler {
	return &SystemParameterHandler{useCase}
}

func (handler *SystemParameterHandler) Hello(c echo.Context) error {
	hello, err := handler.useCase.Hello(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Hello, World! "+hello)
}

func (handler *SystemParameterHandler) Create(c echo.Context) error {
	request := new(dto.SystemParameterRequest)
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = validator.ReqBody(c, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, request)
}

func (handler *SystemParameterHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var request dto.SystemParameterRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(id)
	return c.JSON(http.StatusCreated, request)

}

func (handler *SystemParameterHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

func (handler *SystemParameterHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}
