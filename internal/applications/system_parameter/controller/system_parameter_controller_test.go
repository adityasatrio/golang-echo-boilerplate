package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
	mock_service "myapp/mocks/system_parameter/service"
	"myapp/mocks/test_helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSystemParameterController_Create(t *testing.T) {

	e := test_helper.InitEchoTest(t)

	// CreateTx a new request with sample data
	data := `{"Key":"test_param","Value":"1234"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// CreateTx a new mock service
	mockService := &mock_service.SystemParameterService{}

	// Initialize a new controller
	controller := NewSystemParameterController(mockService)

	dtoCreate := &dto.SystemParameterCreateRequest{
		Key:   "test_param",
		Value: "1234",
	}

	// Test Create function
	mockService.On("Create", req.Context(), dtoCreate).Return(&ent.SystemParameter{}, nil)
	if assert.NoError(t, controller.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestSystemParameterController_Update(t *testing.T) {

	e := test_helper.InitEchoTest(t)

	// CreateTx a new request with sample data
	data := `{"Key":"test_param","Value":"1234"}`
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
		Key:   "test_param",
		Value: "1234",
	}

	// Test UpdateTx function
	mockService.On("Update", req.Context(), 1, dtoUpdate).Return(&ent.SystemParameter{}, nil)
	if assert.NoError(t, controller.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSystemParameterController_Delete(t *testing.T) {
	// Initialize a new echo router
	e := echo.New()

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
	mockService.On("Delete", mock.Anything, mock.Anything).Return(&dto.SystemParameterResponse{}, nil)
	if assert.NoError(t, controller.Delete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

/*
func (caseStub *SystemParameterCaseStub) Hello(ctx context.Context) (string, error) {
	//TODO implement me
	return "hello from stub", nil
}

func (caseStub *SystemParameterCaseStub) CreateSystemParameter(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}

func (caseStub *SystemParameterCaseStub) UpdateSystemParameter(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}

func (caseStub *SystemParameterCaseStub) DeleteSystemParameter(ctx context.Context) error {
	//TODO implement me
	return nil
}

func (caseStub *SystemParameterCaseStub) GetSystemParameterById(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}

func (caseStub *SystemParameterCaseStub) GetSystemParameterAll(ctx context.Context) ([]*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}
*/
/*func TestHello(t *testing.T) {
	// SetupRouteHandler
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/system-parameter", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	//setup stub
	useCaseStub := &SystemParameterCaseStub{}
	h := NewSystemParameterController(useCaseStub)

	//test global_handler
	expected := "\"Hello, World! hello from stub\"\n"
	err := h.Hello(ctx)

	//assert global_handler result
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.ErrorCode)
		assert.Equal(t, expected, recorder.Body.String())
	}

	//assert use case stub
	actualResult, err := h.service.Hello(ctx.Request().Context())

	//assert use case stub
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.ErrorCode)
		assert.Equal(t, "hello from stub", actualResult)
	}

}*/
