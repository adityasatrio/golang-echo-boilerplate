package controller

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/apputils"
	mock_service "myapp/mocks/system_parameter/service"
	"myapp/mocks/test_helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSystemParameterController_Create(t *testing.T) {

	e := test_helper.InitEchoTest(t)

	// Create a new request with sample data
	data := `{"Key":"key1","Value":"value1"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Create a new mock service
	mockService := &mock_service.SystemParameterService{}

	// Initialize a new controller
	controller := NewSystemParameterController(mockService)

	dtoCreate := &dto.SystemParameterCreateRequest{
		Key:   "key1",
		Value: "value1",
	}

	// Test Create function
	mockService.On("Create", req.Context(), dtoCreate).Return(&ent.SystemParameter{Key: "key1", Value: "value1"}, nil)
	if assert.NoError(t, controller.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		dataKey, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Key")
		assert.Equal(t, "key1", dataKey)

		dataValue, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Value")
		assert.Equal(t, "value1", dataValue)
	}
}

func TestSystemParameterController_Update(t *testing.T) {

	e := test_helper.InitEchoTest(t)

	// CreateTx a new request with sample data
	data := `{"Key":"key1","Value":"value1"}`
	req := httptest.NewRequest(http.MethodPut, "/1", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// CreateTx a new mock service
	mockService := &mock_service.SystemParameterService{}

	// Initialize a new controller
	controller := NewSystemParameterController(mockService)

	dtoUpdate := &dto.SystemParameterUpdateRequest{
		Key:   "key1",
		Value: "value1",
	}

	// Test UpdateTx function
	mockService.On("Update", req.Context(), 1, dtoUpdate).Return(&ent.SystemParameter{Key: "key1", Value: "value1"}, nil)
	if assert.NoError(t, controller.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		dataKey, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Key")
		assert.Equal(t, "key1", dataKey)

		dataValue, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Value")
		assert.Equal(t, "value1", dataValue)
	}
}

func TestSystemParameterController_Delete(t *testing.T) {
	e := test_helper.InitEchoTest(t)

	// CreateTx a new request
	req := httptest.NewRequest(http.MethodDelete, "/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// CreateTx a new mock service
	mockService := &mock_service.SystemParameterService{}

	// Initialize a new controller
	controller := NewSystemParameterController(mockService)

	// Test Delete function
	mockService.On("Delete", req.Context(), 1).Return(&ent.SystemParameter{Key: "key1", Value: "value1"}, nil)
	if assert.NoError(t, controller.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		dataKey, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Key")
		assert.Equal(t, "key1", dataKey)

		dataValue, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Value")
		assert.Equal(t, "value1", dataValue)
	}
}

func TestSystemParameterController_GetByID(t *testing.T) {
	e := test_helper.InitEchoTest(t)

	// CreateTx a new request
	req := httptest.NewRequest(http.MethodGet, "/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// CreateTx a new mock service
	mockService := &mock_service.SystemParameterService{}

	// Initialize a new controller
	controller := NewSystemParameterController(mockService)

	// Test Delete function
	mockService.On("GetById", req.Context(), 1).Return(&ent.SystemParameter{Key: "key1", Value: "value1"}, nil)
	if assert.NoError(t, controller.GetById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		dataKey, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Key")
		assert.Equal(t, "key1", dataKey)

		dataValue, _ := apputils.GetFieldBytes(rec.Body.Bytes(), "data.Value")
		assert.Equal(t, "value1", dataValue)
	}
}

func TestSystemParameterController_GetAll(t *testing.T) {
	e := test_helper.InitEchoTest(t)

	// CreateTx a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//c.SetParamNames("id")
	//c.SetParamValues("1")

	// CreateTx a new mock service
	mockService := &mock_service.SystemParameterService{}

	// Initialize a new controller
	controller := NewSystemParameterController(mockService)

	// Test Delete function
	mockService.On("GetAll", req.Context()).Return([]*ent.SystemParameter{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
		{Key: "key3", Value: "value3"},
	}, nil)

	if assert.NoError(t, controller.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Assert the "data" field is an empty array
		data, _ := apputils.GetResultBytes(rec.Body.Bytes(), "data")
		assert.True(t, data.IsArray())
		assert.Equal(t, 3, len(data.Array()))

		data.ForEach(func(_, value gjson.Result) bool {
			key := value.Get("Key").String()
			val := value.Get("Value").String()
			// Perform assertions on each element
			assert.NotNil(t, key)
			assert.NotNil(t, val)
			// Continue iterating
			return true
		})
	}
}
