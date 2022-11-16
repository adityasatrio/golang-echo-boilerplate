package handler

import (
	"github.com/labstack/echo/v4"
	"myapp/modules/system_parameter/usecase"
	"net/http"
)

type SystemParameterHandler struct {
	useCase usecase.SystemParameterCase
}

func New(useCase usecase.SystemParameterCase) *SystemParameterHandler {
	return &SystemParameterHandler{useCase}
}

func (handler *SystemParameterHandler) CreateSystemParameter(c echo.Context) error {
	hello, err := handler.useCase.Hello(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Hello, World!"+hello)
}
