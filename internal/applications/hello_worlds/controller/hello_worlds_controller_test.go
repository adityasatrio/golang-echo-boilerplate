package controller

import (
	"myapp/internal/applications/hello_worlds/service"
	"testing"
)

type MockHelloWorldService struct {
	service.HelloWorldsService

	//mockNameFunction - func - return value of service
	MockHello func() (string, error)
}

func (mock *MockHelloWorldService) Hello() (string, error) {
	return mock.MockHello()
}

/* enable when already use logger
func setUp_author_test() func() {
	logger.New()

	return func() {
		logger.Sync()
		logger.Delete()
	}
}*/

func TestHello(t *testing.T) {
	//	defer setUp_author_test()()

	//mockService := &service.
	//setup
	//e := echo.New()
	//request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	//request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// := httptest.NewRecorder()
	//c := e.NewContext(request, record)

	//mockService := &service.MockHelloWorldsService{}
	//mockService.On("Hello", context.Background(), "hello from controller -", "").Return("Hello World", nil)

	//controller := &HelloWorldsController{}

}
