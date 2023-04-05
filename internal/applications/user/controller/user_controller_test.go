package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"myapp/config/validator"
	"myapp/ent"
	"myapp/internal/applications/user/dto"
	mockService "myapp/mocks/user/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var userService = new(mockService.UserService)
var controller = NewUserController(userService)

func TestUserController_Create(t *testing.T) {

	e := echo.New()
	validator.SetupValidator(e)

	// Create a request.
	reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "tentangAnak123!", "role_id" : 1}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	data := buildUserRequest(reqBody)

	// Create a response recorder.
	res := httptest.NewRecorder()

	// Create a mock response.
	mockRes := buildUserResponse(reqBody)

	userService.On("Create", req.Context(), &data).Return(mockRes, nil)

	// Call the Create method of the UserController.
	c := e.NewContext(req, res)
	err := controller.Create(c)

	// Check the http response:
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.Code)

	// Check the data response:
	resBody := mockRes
	err = json.NewDecoder(res.Body).Decode(resBody)
	assert.NoError(t, err)
	assert.Equal(t, mockRes, resBody)

	userService.AssertExpectations(t)
}

func buildUserRequest(reqJson string) dto.UserRequest {
	data := dto.UserRequest{}
	errJson := json.Unmarshal([]byte(reqJson), &data)

	if errJson != nil {
		log.Println(errJson)
		return data
	}

	return data
}

func buildUserResponse(reqJson string) *ent.User {
	data := ent.User{}
	errJson := json.Unmarshal([]byte(reqJson), &data)

	if errJson != nil {
		log.Println(errJson)
		return &data
	}

	return &data
}
