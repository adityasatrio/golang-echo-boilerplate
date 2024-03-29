package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/internal/helper"
	mock_service "myapp/mocks/applications/health/service"
	"myapp/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHealthController(t *testing.T) {
	mockService := &mock_service.HealthService{}

	// Call the function being tested.
	controller := NewHealthController(mockService)

	// Check that the service field of the controller has the expected value.
	if controller.service != mockService {
		t.Errorf("Expected service to be %v, but got %v", mockService, controller.service)
	}
}

func TestHealthDependency_Success(t *testing.T) {

	e := test.InitEchoTest(t)

	request := httptest.NewRequest(http.MethodGet, "/skypiea/health?flag=dependency", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	healthCheck := map[string]string{}
	healthCheck["ctx_status"] = "UP"
	healthCheck["ctx_name"] = "echo"
	healthCheck["db_status"] = "UP"
	healthCheck["db_name"] = "mysql"
	healthCheck["cache_status"] = "UP"
	healthCheck["cache_name"] = "cache"
	healthCheck["final_msg"] = "hello from controller layer " + "hello from repository layer " + "hello from query parameter"

	mockService := &mock_service.HealthService{}
	mockService.On("Health", request.Context(), "hello from controller layer ").
		Return(healthCheck, nil)

	controller := &HealthController{mockService}

	// Assertions
	if assert.NoError(t, controller.Health(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		dataKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.status")
		assert.Equal(t, "UP", dataKey)

		dataMessageKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.message")
		assert.Equal(t, "hello from controller layer "+"hello from repository layer "+"hello from query parameter", dataMessageKey)

		dataComponentKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.components")
		assert.NotNil(t, dataComponentKey)
	}
}

func TestHealthDependency_Failed(t *testing.T) {

	e := test.InitEchoTest(t)

	request := httptest.NewRequest(http.MethodGet, "/skypiea/health?flag=dependency", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	healthCheckMock := map[string]string{}
	healthCheckMock["ctx_status"] = "UP"
	healthCheckMock["ctx_name"] = "echo"
	healthCheckMock["db_status"] = "DOWN"
	healthCheckMock["db_name"] = "mysql"
	healthCheckMock["cache_status"] = "DOWN"
	healthCheckMock["cache_name"] = "cache"
	healthCheckMock["final_msg"] = "hello from controller layer " + "hello from repository layer " + "hello from query parameter"

	mockService := &mock_service.HealthService{}
	mockService.On("Health", request.Context(), "hello from controller layer ").
		Return(healthCheckMock, errors.New("system down"))

	controller := &HealthController{mockService}

	// Assertions
	if assert.NoError(t, controller.Health(c)) {
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		dataKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.status")
		assert.Equal(t, "UP", dataKey)

		dataDb, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.components.db.status")
		assert.Equal(t, "DOWN", dataDb)

		dataCache, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.components.cache.status")
		assert.Equal(t, "DOWN", dataCache)

		dataMessageKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.message")
		assert.Equal(t, "hello from controller layer "+"hello from repository layer "+"hello from query parameter", dataMessageKey)

		dataComponentKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.components")
		assert.NotNil(t, dataComponentKey)
	}
}

func TestHealthNoDependency(t *testing.T) {

	e := test.InitEchoTest(t)

	request := httptest.NewRequest(http.MethodGet, "/skypiea/health", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	mockService := &mock_service.HealthService{}
	controller := &HealthController{mockService}

	// Assertions
	if assert.NoError(t, controller.Health(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		dataKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.status")
		assert.Equal(t, "UP", dataKey)

		dataMessageKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.message")
		assert.Equal(t, "", dataMessageKey)

		dataComponentKey, _ := helper.GetFieldBytes(recorder.Body.Bytes(), "data.components.ctx.status")
		assert.Equal(t, nil, dataComponentKey)
	}
}
