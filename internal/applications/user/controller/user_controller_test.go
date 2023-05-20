package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/configs/validator"
	"myapp/ent"
	"myapp/internal/applications/user/dto"
	"myapp/internal/apputils"
	mockService "myapp/mocks/user/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserController_Create(t *testing.T) {

	userMocks := []struct {
		name              string
		reqBodyStrJson    string
		reqBodyJson       dto.UserRequest
		resBodyJson       dto.UserResponse
		userServiceResult ent.User
		scenario          bool
		errors            error
	}{
		{
			name:           "Create_success",
			reqBodyStrJson: `{"name": "testing name", "email": "testing@email.id", "password": "Password123!", "role_id" : 1}`,
			reqBodyJson: dto.UserRequest{
				RoleId:   1,
				Name:     "testing name",
				Email:    "testing@email.id",
				Password: "Password123!",
			},
			resBodyJson: dto.UserResponse{
				Id:       1,
				RoleId:   1,
				Name:     "testing name",
				Email:    "testing@email.id",
				Password: "Password123!",
				Avatar:   "xx/ava.png",
			},
			userServiceResult: ent.User{
				ID:       1,
				RoleID:   1,
				Name:     "testing name",
				Email:    "testing@email.id",
				Password: "Password123!",
				Avatar:   "xx/ava.png",
			},
			scenario: true,
			errors:   nil,
		}, {
			name:           "Create_failed_whenCallService",
			reqBodyStrJson: `{"name": "testing name", "email": "testing@email.id", "password": "Password123!", "role_id" : 1}`,
			reqBodyJson: dto.UserRequest{
				RoleId:   1,
				Name:     "testing name",
				Email:    "testing@email.id",
				Password: "Password123!",
			},
			resBodyJson: dto.UserResponse{
				Id:       1,
				RoleId:   1,
				Name:     "testing name",
				Email:    "testing@email.id",
				Password: "Password123!",
				Avatar:   "xx/ava.png",
			},
			userServiceResult: ent.User{
				ID:       1,
				RoleID:   1,
				Name:     "testing name",
				Email:    "testing@email.id",
				Password: "Password123!",
				Avatar:   "xx/ava.png",
			},
			scenario: false,
			errors:   errors.New("got failed when create user"),
		},
	}

	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {

			userService := new(mockService.UserService)
			controller := NewUserController(userService)
			e := echo.New()
			validator.SetupValidator(e)

			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userMock.reqBodyStrJson))
			req.Header.Set("Content-Type", "application/json")

			userService.On("Create", req.Context(), &userMock.reqBodyJson).Return(&userMock.userServiceResult, userMock.errors)

			// Call the CreateTx method of the UserController.
			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := controller.Create(c)

			if userMock.scenario {

				// Check the http response:
				assert.NoError(t, err)
				assert.Equal(t, http.StatusCreated, res.Code)

				//code, msg, errCode := validator.MapperErrorCode(err)
				//assert.Equal(t, 400, code)
				//assert.NotNil(t, msg)
				//assert.NotNil(t, errCode)

				resName, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.name")
				assert.Equal(t, "testing name", resName)

				resEmail, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.email")
				assert.Equal(t, "testing@email.id", resEmail)

				resPassword, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.password")
				assert.Equal(t, "Password123!", resPassword)

			} else {

				// Check the http response:
				assert.Error(t, err)
				assert.Equal(t, http.StatusOK, res.Code)

				resName, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.name")
				assert.Nil(t, resName)

				resEmail, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.email")
				assert.Nil(t, resEmail)

				resPassword, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.Password")
				assert.Nil(t, resPassword)

			}

		})
	}

	//t.Run("Create_failed_validation", func(t *testing.T) {
	//
	//	userService := new(mockService.UserService)
	//	controller := NewUserController(userService)
	//
	//	e := echo.New()
	//	validator.SetupValidator(e)
	//	validator.SetupGlobalHttpUnhandleErrors(e)
	//
	//	//reqBody := `{"name": "testing name", "email": "invalid_email", "password": "password", "role_id" : 1}`
	//	reqBody := `{}`
	//	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	//	req.Header.Set("Content-Type", "application/json")
	//
	//	// Call the  method of the UserController.
	//	res := httptest.NewRecorder()
	//	c := e.NewContext(req, res)
	//
	//	err := controller.Create(c)
	//
	//	// Check the http response:
	//	//TODO : should be have err and handle error with #response_json_builder.go generic error mapper
	//	assert.NoError(t, err) //not error because when invalid we directly return bad request
	//	assert.Equal(t, http.StatusBadRequest, res.Code)
	//
	//	resData, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data")
	//	assert.NotNil(t, resData)
	//
	//	resErr, _ := apputils.GetFieldBytes(res.Body.Bytes(), "error")
	//	assert.NotNil(t, resErr)
	//
	//})

}

func TestUserController_Update(t *testing.T) {

	reqBodyJson := dto.UserRequest{
		RoleId:   1,
		Name:     "testing name",
		Email:    "testing@email.id",
		Password: "Password123!",
	}

	userServiceResult := ent.User{
		ID:       1,
		RoleID:   1,
		Name:     "testing name",
		Email:    "testing@email.id",
		Password: "Password123!",
		Avatar:   "xx/ava.png",
	}

	t.Run("Update_success", func(t *testing.T) {

		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// CreateTx a request.
		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "Password123!", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		userService.On("Update", req.Context(), uint64(1), &reqBodyJson).Return(&userServiceResult, nil)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.Update(c)

		// Check the http response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		resName, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.name")
		assert.Equal(t, "testing name", resName)

		resEmail, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.email")
		assert.Equal(t, "testing@email.id", resEmail)

		resPassword, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.password")
		assert.Equal(t, "Password123!", resPassword)
	})

	t.Run("Update_failed", func(t *testing.T) {
		userService := new(mockService.UserService)
		controller := NewUserController(userService)
		e := echo.New()
		validator.SetupValidator(e)

		// CreateTx a request.
		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "Password123!", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		errorMessage := errors.New("failed got user")
		userService.On("Update", req.Context(), uint64(1), &reqBodyJson).Return(nil, errorMessage)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.Update(c)

		// Check the http response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.Code)

		resData, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data")
		assert.NotNil(t, resData)

		resErr, _ := apputils.GetFieldBytes(res.Body.Bytes(), "error")
		assert.NotNil(t, resErr)
	})

	//t.Run("Update_failed_validation", func(t *testing.T) {
	//	userService := new(mockService.UserService)
	//	controller := NewUserController(userService)
	//	e := echo.New()
	//	validator.SetupValidator(e)
	//
	//	// CreateTx a request.
	//	reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
	//	req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
	//	req.Header.Set("Content-Type", "application/json")
	//
	//	// CreateTx a mock response.
	//	mockRes := buildUserResponse(reqBody)
	//
	//	// Call the CreateTx method of the UserController.
	//	res := httptest.NewRecorder()
	//	c := e.NewContext(req, res)
	//	c.SetParamNames("id")
	//	c.SetParamValues("1")
	//	err := controller.Update(c)
	//
	//	// Check the http response:
	//	assert.Error(t, err)
	//	assert.Equal(t, http.StatusOK, res.Code)
	//
	//	// Check the data response:
	//	resBody := mockRes
	//	err = json.NewDecoder(res.Body).Decode(resBody)
	//	assert.Error(t, err)
	//	assert.Equal(t, mockRes, resBody)
	//
	//	userService.AssertExpectations(t)
	//})
}

//func TestUserController_Delete(t *testing.T) {
//	t.Run("Delete_success", func(t *testing.T) {
//
//		userService := new(mockService.UserService)
//		controller := NewUserController(userService)
//		e := echo.New()
//		validator.SetupValidator(e)
//
//		// CreateTx a request.
//		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(""))
//		req.Header.Set("Content-Type", "application/json")
//
//		// CreateTx a mock response.
//		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
//		mockRes := buildUserResponse(reqBody)
//
//		userService.On("Delete", req.Context(), uint64(1)).Return(mockRes, nil)
//
//		// Call the CreateTx method of the UserController.
//		res := httptest.NewRecorder()
//		c := e.NewContext(req, res)
//		c.SetParamNames("id")
//		c.SetParamValues("1")
//		err := controller.Delete(c)
//
//		// Check the http response:
//		assert.NoError(t, err)
//		assert.Equal(t, http.StatusOK, res.Code)
//
//		// Check the data response:
//		resBody := mockRes
//		err = json.NewDecoder(res.Body).Decode(resBody)
//		assert.NoError(t, err)
//		assert.Equal(t, mockRes, resBody)
//
//		userService.AssertExpectations(t)
//	})
//
//	t.Run("Delete_failed", func(t *testing.T) {
//
//		userService := new(mockService.UserService)
//		controller := NewUserController(userService)
//		e := echo.New()
//		validator.SetupValidator(e)
//
//		// CreateTx a request.
//		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(""))
//		req.Header.Set("Content-Type", "application/json")
//
//		// CreateTx a mock response.
//		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
//		mockRes := buildUserResponse(reqBody)
//
//		errorMessage := errors.New("failed delete user")
//		userService.On("Delete", req.Context(), uint64(1)).Return(mockRes, errorMessage)
//
//		// Call the CreateTx method of the UserController.
//		res := httptest.NewRecorder()
//		c := e.NewContext(req, res)
//		c.SetParamNames("id")
//		c.SetParamValues("1")
//		err := controller.Delete(c)
//
//		// Check the http response:
//		assert.Error(t, err)
//		assert.Equal(t, http.StatusOK, res.Code)
//
//		// Check the data response:
//		resBody := mockRes
//		err = json.NewDecoder(res.Body).Decode(resBody)
//		assert.Error(t, err)
//		assert.Equal(t, mockRes, resBody)
//
//		userService.AssertExpectations(t)
//	})
//
//}

//func TestUserController_GetById(t *testing.T) {
//
//	t.Run("Get_success", func(t *testing.T) {
//
//		userService := new(mockService.UserService)
//		controller := NewUserController(userService)
//		e := echo.New()
//		validator.SetupValidator(e)
//
//		// CreateTx a request.
//		req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
//		req.Header.Set("Content-Type", "application/json")
//
//		// CreateTx a mock response.
//		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
//		mockRes := buildUserResponse(reqBody)
//
//		userService.On("GetById", req.Context(), uint64(1)).Return(mockRes, nil)
//
//		// Call the CreateTx method of the UserController.
//		res := httptest.NewRecorder()
//		c := e.NewContext(req, res)
//		c.SetParamNames("id")
//		c.SetParamValues("1")
//		err := controller.GetById(c)
//
//		// Check the http response:
//		assert.NoError(t, err)
//		assert.Equal(t, http.StatusOK, res.Code)
//
//		// Check the data response:
//		resBody := mockRes
//		err = json.NewDecoder(res.Body).Decode(resBody)
//		assert.NoError(t, err)
//		assert.Equal(t, mockRes, resBody)
//
//		userService.AssertExpectations(t)
//	})
//
//	t.Run("Get_failed", func(t *testing.T) {
//
//		userService := new(mockService.UserService)
//		controller := NewUserController(userService)
//		e := echo.New()
//		validator.SetupValidator(e)
//
//		// CreateTx a request.
//		req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
//		req.Header.Set("Content-Type", "application/json")
//
//		// CreateTx a mock response.
//		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
//		mockRes := buildUserResponse(reqBody)
//
//		errorMessage := errors.New("failed get user")
//		userService.On("GetById", req.Context(), uint64(1)).Return(mockRes, errorMessage)
//
//		// Call the CreateTx method of the UserController.
//		res := httptest.NewRecorder()
//		c := e.NewContext(req, res)
//		c.SetParamNames("id")
//		c.SetParamValues("1")
//		err := controller.GetById(c)
//
//		// Check the http response:
//		assert.Error(t, err)
//		assert.Equal(t, http.StatusOK, res.Code)
//
//		// Check the data response:
//		resBody := mockRes
//		err = json.NewDecoder(res.Body).Decode(resBody)
//		assert.Error(t, err)
//		assert.Equal(t, mockRes, resBody)
//
//		userService.AssertExpectations(t)
//	})
//}

//func TestUserController_GetAll(t *testing.T) {
//
//	t.Run("Get_All_success", func(t *testing.T) {
//
//		userService := new(mockService.UserService)
//		controller := NewUserController(userService)
//		e := echo.New()
//		validator.SetupValidator(e)
//
//		// CreateTx a request.
//		req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
//		req.Header.Set("Content-Type", "application/json")
//
//		// CreateTx a mock response.
//		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
//		mockRes := buildUserResponse(reqBody)
//
//		mockListUser := make([]*ent.User, 0)
//		mockListUser = append(mockListUser, mockRes)
//
//		userService.On("GetAll", req.Context()).Return(mockListUser, nil)
//
//		// Call the CreateTx method of the UserController.
//		res := httptest.NewRecorder()
//		c := e.NewContext(req, res)
//		err := controller.GetAll(c)
//
//		// Check the http response:
//		assert.NoError(t, err)
//		assert.Equal(t, http.StatusOK, res.Code)
//
//		// Check the data response:
//		resBody := mockRes
//		err = json.NewDecoder(res.Body).Decode(resBody)
//		assert.NoError(t, err)
//		assert.Equal(t, mockRes, resBody)
//
//		userService.AssertExpectations(t)
//	})
//
//	t.Run("Get_All_failed", func(t *testing.T) {
//
//		userService := new(mockService.UserService)
//		controller := NewUserController(userService)
//		e := echo.New()
//		validator.SetupValidator(e)
//
//		// CreateTx a request.
//		req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
//		req.Header.Set("Content-Type", "application/json")
//
//		// CreateTx a mock response.
//		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "password", "role_id" : 1}`
//		mockRes := buildUserResponse(reqBody)
//
//		errorMessage := errors.New("failed get user")
//		userService.On("GetAll", req.Context()).Return(nil, errorMessage)
//
//		// Call the CreateTx method of the UserController.
//		res := httptest.NewRecorder()
//		c := e.NewContext(req, res)
//		err := controller.GetAll(c)
//
//		// Check the http response:
//		assert.Error(t, err)
//		assert.Equal(t, http.StatusOK, res.Code)
//
//		// Check the data response:
//		resBody := mockRes
//		err = json.NewDecoder(res.Body).Decode(resBody)
//		assert.Error(t, err)
//		assert.Equal(t, mockRes, resBody)
//
//		userService.AssertExpectations(t)
//	})
//}

//
//func buildUserRequest(reqJson string) dto.UserRequest {
//	data := dto.UserRequest{}
//	errJson := json.Unmarshal([]byte(reqJson), &data)
//
//	if errJson != nil {
//		log.Println(errJson)
//		return data
//	}
//
//	return data
//}
//
//func buildUserResponse(reqJson string) *ent.User {
//	data := ent.User{}
//	errJson := json.Unmarshal([]byte(reqJson), &data)
//
//	if errJson != nil {
//		log.Println(errJson)
//		return &data
//	}
//
//	return &data
//}
