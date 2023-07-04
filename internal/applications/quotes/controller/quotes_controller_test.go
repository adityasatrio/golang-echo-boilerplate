package controller

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/configs/validator"
	"myapp/exceptions"
	"myapp/internal/applications/quotes/dto"
	"myapp/internal/helper"
	mock_service "myapp/mocks/quotes/service"
	"myapp/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestQuotesController_Get(t *testing.T) {
	e := test.InitEchoTest(t)

	req := httptest.NewRequest(http.MethodGet, "/quotes?query1=value1&query2=value2", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := &mock_service.QuotesService{}
	controller := NewQuotesController(mockService)

	// Test function
	mockQueryParameter := map[string]string{
		"query1": c.QueryParam("query1"),
		"query2": c.QueryParam("query2"),
	}

	mockService.On("GetQuotes", req.Context(), mockQueryParameter).Return(&dto.QuoteApiResponse{Author: "key1", Quote: "value1"}, nil)
	if assert.NoError(t, controller.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		dataKey, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.author")
		assert.Equal(t, "key1", dataKey)

		dataValue, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.quote")
		assert.Equal(t, "value1", dataValue)

		dataCustom, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.custom")
		assert.Equal(t, "beyond limit", dataCustom)
	}
}

func TestQuotesController_GetError(t *testing.T) {
	e := test.InitEchoTest(t)

	req := httptest.NewRequest(http.MethodGet, "/quotes?query1=value1&query2=value2", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := &mock_service.QuotesService{}
	controller := NewQuotesController(mockService)

	errorMessage := errors.New("get data failed")
	errEx := exceptions.NewBusinessLogicError(exceptions.DataGetFailed, errorMessage)

	// Test function
	mockQueryParameter := map[string]string{
		"query1": c.QueryParam("query1"),
		"query2": c.QueryParam("query2"),
	}
	mockService.On("GetQuotes", req.Context(), mockQueryParameter).Return(nil, errEx)
	controllerErr := controller.Get(c)

	assert.Error(t, controllerErr)

	errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(controllerErr)
	assert.Equal(t, http.StatusUnprocessableEntity, errHttpCode)
	assert.Equal(t, exceptions.DataGetFailed, errBusinessCode)
	assert.Equal(t, 10006, errBusinessCode)
	assert.Equal(t, "get data failed", errMsg)
	assert.Nil(t, errOut)
}

func TestQuotesController_Post(t *testing.T) {
	e := test.InitEchoTest(t)

	reqBody := `{"name": "testing", "author": "testing", "quote": "testing", "custom" : "testing"}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := &mock_service.QuotesService{}
	controller := NewQuotesController(mockService)

	mockReqBody := &dto.QuoteRequest{
		Name:   "testing",
		Author: "testing",
		Quote:  "testing",
		Custom: "testing",
	}

	mockService.On("PostQuotes", req.Context(), mockReqBody).Return(&dto.QuoteApiResponse{Author: "key1", Quote: "value1"}, nil)
	if assert.NoError(t, controller.Post(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		dataKey, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.author")
		assert.Equal(t, "key1", dataKey)

		dataValue, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.quote")
		assert.Equal(t, "value1", dataValue)

		dataCustom, _ := helper.GetFieldBytes(rec.Body.Bytes(), "data.custom")
		assert.Equal(t, "beyond limit", dataCustom)
	}
}

func TestQuotesController_PostFailed(t *testing.T) {
	e := test.InitEchoTest(t)

	reqBody := `{"name": "testing", "author": "testing", "quote": "testing", "custom" : "testing"}`
	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := &mock_service.QuotesService{}
	controller := NewQuotesController(mockService)

	errorMessage := errors.New("create data failed")
	errEx := exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, errorMessage)

	mockReqBody := &dto.QuoteRequest{
		Name:   "testing",
		Author: "testing",
		Quote:  "testing",
		Custom: "testing",
	}

	mockService.On("PostQuotes", req.Context(), mockReqBody).Return(nil, errEx)
	controllerErr := controller.Post(c)

	assert.Error(t, controllerErr)

	errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(controllerErr)
	assert.Equal(t, http.StatusUnprocessableEntity, errHttpCode)
	assert.Equal(t, exceptions.DataCreateFailed, errBusinessCode)
	assert.Equal(t, 10003, errBusinessCode)
	assert.Equal(t, "create data failed", errMsg)
	assert.Nil(t, errOut)
}
