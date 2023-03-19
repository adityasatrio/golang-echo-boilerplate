package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock_service "myapp/mocks/hello_worlds/service"
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

	//mockCtx := mocks.NewMockContext(t)
	mockService := &mock_service.HelloWorldsService{}
	mockService.On("Hello", request.Context(), "hello from controller -", "").Return("success", nil)

	controller := &HelloWorldsController{mockService}

	// Assertions
	if assert.NoError(t, controller.Hello(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		// Get the response body as a string
		responseString := recorder.Body.String()

		// Convert the response string to JSON
		var jsonResponse map[string]interface{}
		err := json.Unmarshal([]byte(responseString), &jsonResponse)
		if err != nil {
			t.Errorf("Error unmarshalling JSON: %v", err)
		}

		//sample response
		//"{\"code\":200,\"message\":\"OK\",\"data\":\"success\",\"error\":\"\",\"serverTime\":\"Sun, 19 Mar 2023 19:20:57 WIB\"}\n"
		assert.Equal(t, "success", jsonResponse["data"])

	}
}
