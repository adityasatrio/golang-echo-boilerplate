package controller

import (
	"encoding/json"
	"errors"
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

func TestUserController_Create(t *testing.T) {

	userMocks := []struct {
		name        string
		reqBodyJson string
		scenario    bool
		errors      error
	}{
		{
			name:        "Create_success",
			reqBodyJson: `{"name": "testing name", "email": "testing@tentanganak.id", "password": "tentangAnak123!", "role_id" : 1}`,
			scenario:    true,
			errors:      nil,
		}, {
			name:        "Create_failed_whenCallService",
			reqBodyJson: `{"name": "testing name", "email": "testing@tentanganak.id", "password": "tentangAnak123!", "role_id" : 1}`,
			scenario:    false,
			errors:      errors.New("got failed when create user"),
		},
	}

	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {

			userService := new(mockService.UserService)
			controller := NewUserController(userService)
			e := echo.New()
			validator.SetupValidator(e)

			// Create a request.
			reqBody := userMock.reqBodyJson
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			// Create a mock response.
			mockReq := buildUserRequest(reqBody)

			// Create a mock response.
			mockRes := buildUserResponse(reqBody)

			userService.On("Create", req.Context(), &mockReq).Return(mockRes, userMock.errors)

			// Call the Create method of the UserController.
			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := controller.Create(c)

			if userMock.scenario {

				// Check the http response:
				assert.NoError(t, err)
				assert.Equal(t, http.StatusCreated, res.Code)

				// Check the data response:
				resBody := mockRes
				err = json.NewDecoder(res.Body).Decode(resBody)
				assert.NoError(t, err)
				assert.Equal(t, mockRes, resBody)
				userService.AssertExpectations(t)

			} else {

				// Check the http response:
				assert.Error(t, err)
				assert.Equal(t, http.StatusOK, res.Code)

				// Check the data response:
				err = json.NewDecoder(res.Body).Decode(mockRes)
				assert.Error(t, err)
				userService.AssertExpectations(t)
			}

		})
	}

	t.Run("Create_failed_validation", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		mockRes := buildUserResponse(reqBody)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := controller.Create(c)

		// Check the http response:
		assert.Error(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.Error(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

}

func TestUserController_Update(t *testing.T) {
	t.Run("Update_success", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "tentangAnak123!", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		mockReq := buildUserRequest(reqBody)

		// Create a mock response.
		mockRes := buildUserResponse(reqBody)

		userService.On("Update", req.Context(), uint64(1), &mockReq).Return(mockRes, nil)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.Update(c)

		// Check the http response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.NoError(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

	t.Run("Update_failed", func(t *testing.T) {
		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "tentangAnak123!", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		mockReq := buildUserRequest(reqBody)

		// Create a mock response.
		mockRes := buildUserResponse(reqBody)

		errorMessage := errors.New("failed got user")
		userService.On("Update", req.Context(), uint64(1), &mockReq).Return(nil, errorMessage)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.Update(c)

		// Check the http response:
		assert.Error(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.Error(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

	t.Run("Update_failed_validation", func(t *testing.T) {
		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		mockRes := buildUserResponse(reqBody)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.Update(c)

		// Check the http response:
		assert.Error(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.Error(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})
}

func TestUserController_Delete(t *testing.T) {
	t.Run("Delete_success", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		mockRes := buildUserResponse(reqBody)

		userService.On("Delete", req.Context(), uint64(1)).Return(mockRes, nil)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.Delete(c)

		// Check the http response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.NoError(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

	t.Run("Delete_failed", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		mockRes := buildUserResponse(reqBody)

		errorMessage := errors.New("failed delete user")
		userService.On("Delete", req.Context(), uint64(1)).Return(mockRes, errorMessage)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.Delete(c)

		// Check the http response:
		assert.Error(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.Error(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

}

func TestUserController_GetById(t *testing.T) {

	t.Run("Get_success", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		mockRes := buildUserResponse(reqBody)

		userService.On("GetById", req.Context(), uint64(1)).Return(mockRes, nil)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.GetById(c)

		// Check the http response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.NoError(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

	t.Run("Get_failed", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		mockRes := buildUserResponse(reqBody)

		errorMessage := errors.New("failed get user")
		userService.On("GetById", req.Context(), uint64(1)).Return(mockRes, errorMessage)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.GetById(c)

		// Check the http response:
		assert.Error(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.Error(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})
}

func TestUserController_GetAll(t *testing.T) {

	t.Run("Get_All_success", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		mockRes := buildUserResponse(reqBody)

		mockListUser := make([]*ent.User, 0)
		mockListUser = append(mockListUser, mockRes)

		userService.On("GetAll", req.Context()).Return(mockListUser, nil)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := controller.GetAll(c)

		// Check the http response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.NoError(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})

	t.Run("Get_All_failed", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// Create a request.
		req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		// Create a mock response.
		reqBody := `{"name": "testing name", "email": "testing@tentanganak.id", "password": "xxxx", "role_id" : 1}`
		mockRes := buildUserResponse(reqBody)

		errorMessage := errors.New("failed get user")
		userService.On("GetAll", req.Context()).Return(nil, errorMessage)

		// Call the Create method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := controller.GetAll(c)

		// Check the http response:
		assert.Error(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		// Check the data response:
		resBody := mockRes
		err = json.NewDecoder(res.Body).Decode(resBody)
		assert.Error(t, err)
		assert.Equal(t, mockRes, resBody)

		userService.AssertExpectations(t)
	})
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
