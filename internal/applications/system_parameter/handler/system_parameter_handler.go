package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/commons/response"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/usecase"
	//"myapp/pkg/validator"
	"net/http"
)

//var err error

type SystemParameterHandler struct {
	useCase usecase.SystemParameterCase //interface
}

/*func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}
*/

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

	//err = validator.ReqBody(c, request)
	err = c.Validate(request)
	fmt.Println("validate", err)
	if err != nil {
		return response.Return(c, http.StatusBadRequest, "failed", err, nil)
		//return c.JSON(http.StatusBadRequest, &err)
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
