package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock_service "myapp/mocks/hello_worlds/service"
	"myapp/mocks/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

//type MockHelloWorldService struct {
//	service.HelloWorldsService
//
//	//mockNameFunction - func - return value of service
//	MockHello func() (string, error)
//}
//
//func (mocks *MockHelloWorldService) Hello() (string, error) {
//	return mocks.MockHello()
//}

/* enable when already use logger
func setUp_author_test() func() {
	logger.New()

	return func() {
		logger.Sync()
		logger.Delete()
	}
}*/

func TestHello(t *testing.T) {

	e := echo.New()

	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	mockCtx := mocks.NewMockContext(t)
	mockService := &mock_service.HelloWorldsService{}

	mockService.On("Hello", mockCtx, "hello", "controller").Return("success", nil)

	controller := &HelloWorldsController{mockService}
	result := controller.Hello(c)

	assert.Equal(t, nil, result)

	//mockService := &service.MockHeloWorldsService{}
	//mockService.On("Hello", context.Background(), "hello from controller -", "").Return("Hello World", nil)

	//controller := &HelloWorldsController{}

}
