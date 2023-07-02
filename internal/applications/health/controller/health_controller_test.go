package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/internal/apputils"
	mock_service "myapp/mocks/health/service"
	"myapp/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHelloWorldsController(t *testing.T) {
	mockService := &mock_service.HealthService{}

	// Call the function being tested.
	controller := NewHealthController(mockService)

	// Check that the service field of the controller has the expected value.
	if controller.service != mockService {
		t.Errorf("Expected service to be %v, but got %v", mockService, controller.service)
	}
}

func TestHealth(t *testing.T) {

	e := test.InitEchoTest(t)

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	healthCheck := map[string]string{}
	healthCheck["ctx_status"] = "UP"
	healthCheck["ctx_name"] = "echo"
	healthCheck["db_status"] = "UP"
	healthCheck["db_name"] = "mysql"
	healthCheck["final_msg"] = "hello from controller layer " + "hello from repository layer " + "hello from query parameter=default"

	mockService := &mock_service.HealthService{}
	mockService.On("Health", request.Context(), "hello from controller layer ", "default").
		Return(healthCheck, nil)

	controller := &HealthController{mockService}

	// Assertions
	if assert.NoError(t, controller.Health(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		dataKey, _ := apputils.GetFieldBytes(recorder.Body.Bytes(), "data.status")
		assert.Equal(t, "UP", dataKey)

	}
}
