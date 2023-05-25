package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/exceptions"
	"myapp/internal/apputils"
	mock_service "myapp/mocks/hello_worlds/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHelloWorldsController(t *testing.T) {
	// CreateTx a mock service object.
	mockService := &mock_service.HelloWorldsService{}

	// Call the function being tested.
	controller := NewHelloWorldsController(mockService)

	// Check that the service field of the controller has the expected value.
	if controller.service != mockService {
		t.Errorf("Expected service to be %v, but got %v", mockService, controller.service)
	}
}

func TestHello(t *testing.T) {

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	mockService := &mock_service.HelloWorldsService{}
	mockService.On("Hello", request.Context(), "hello from controller ->", "").
		Return("success", nil)

	controller := &HelloWorldsController{mockService}

	// Assertions
	if assert.NoError(t, controller.Hello(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		dataKey, _ := apputils.GetFieldBytes(recorder.Body.Bytes(), "data.Message")
		assert.Equal(t, "success", dataKey)

	}
}

func TestHelloErr(t *testing.T) {

	e := echo.New()

	//TODO : this test should be involved global error handle so we can also assert the response code
	//e.HTTPErrorHandler = validator.NewHttpErrorHandler()

	request := httptest.NewRequest(http.MethodGet, "/hello?error=service", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	mockService := &mock_service.HelloWorldsService{}
	serviceErr := exceptions.NewBusinessLogicError(exceptions.EBL10007, errors.New("hello from controller -> hello from s-impl layer ->"))
	mockService.On("Hello", request.Context(), "hello from controller ->", "service").Return("", serviceErr)

	controller := &HelloWorldsController{mockService}

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	// Assertions
	if assert.Errorf(t, serviceErr, "business logic error", controller.Hello(c)) {

		dataKey, _ := apputils.GetFieldBytes(recorder.Body.Bytes(), "data.Message")
		assert.Equal(t, nil, dataKey)
	}
}
