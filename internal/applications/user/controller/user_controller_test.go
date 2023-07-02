package controller

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/configs/validator"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/user/dto"
	"myapp/internal/apputils"
	mock_service "myapp/mocks/user/service"
	"myapp/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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
				ID:       1,
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
				ID:       1,
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

			userService := new(mock_service.UserService)
			controller := NewUserController(userService)
			e := test.InitEchoTest(t)
			validator.SetupValidator(e)

			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userMock.reqBodyStrJson))
			req.Header.Set("Content-Type", "application/json")

			userService.On("Create", req.Context(), &userMock.reqBodyJson).Return(&userMock.userServiceResult, userMock.errors)

			// Call the CreateTx method of the UserController.
			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := controller.Create(c)

			if userMock.scenario {

				// Check the outbound response:
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

				// Check the outbound response:
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

	t.Run("Create_failed_dto_validation", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)

		e := test.InitEchoTest(t)
		validator.SetupValidator(e)
		validator.SetupGlobalHttpUnhandleErrors(e)

		reqBody := `{"name": "testing name", "email": "invalid_email", "password": "password"}`
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// Call the  method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		err := controller.Create(c)

		assert.Error(t, err)

		errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(err)
		assert.Equal(t, http.StatusBadRequest, errHttpCode)
		assert.Equal(t, http.StatusBadRequest, errBusinessCode)
		assert.Equal(t, "RoleId is required", errMsg)
		assert.Nil(t, errOut)

	})

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

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

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

		// Check the outbound response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		resName, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.name")
		assert.Equal(t, "testing name", resName)

		resEmail, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.email")
		assert.Equal(t, "testing@email.id", resEmail)

		resPassword, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.password")
		assert.Equal(t, "Password123!", resPassword)
	})

	t.Run("Update_failed_on_service", func(t *testing.T) {
		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		// CreateTx a request.
		reqBody := `{"name": "testing name", "email": "testing@email.id", "password": "Password123!", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		errorMessage := errors.New("update data failed")
		errEx := exceptions.NewBusinessLogicError(exceptions.EBL10004, errorMessage)

		userService.On("Update", req.Context(), uint64(1), &reqBodyJson).Return(nil, errEx)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.Update(c)

		assert.Error(t, err)

		errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(err)
		assert.Equal(t, http.StatusUnprocessableEntity, errHttpCode)
		assert.Equal(t, exceptions.EBL10004, errBusinessCode)
		assert.Equal(t, "update data failed", errMsg)
		assert.Nil(t, errOut)

	})

	t.Run("Update_failed_dto_validation", func(t *testing.T) {
		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		reqBody := `{"name": "testing name", "email": "testing@email.id", "role_id" : 1}`
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.Update(c)

		assert.Error(t, err)

		errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(err)
		assert.Equal(t, http.StatusBadRequest, errHttpCode)
		assert.Equal(t, http.StatusBadRequest, errBusinessCode)
		assert.Equal(t, "Password is required", errMsg)
		assert.Nil(t, errOut)

	})
}

func TestUserController_Delete(t *testing.T) {
	t.Run("Delete_success", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		// CreateTx a request.
		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		userServiceResult := ent.User{
			ID:        1,
			RoleID:    1,
			Name:      "testing name",
			Email:     "testing@email.id",
			Password:  "Password123!",
			Avatar:    "xx/ava.png",
			DeletedBy: "user",
			DeletedAt: time.Now(),
		}

		userService.On("Delete", req.Context(), uint64(1)).Return(&userServiceResult, nil)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)

		c.SetParamNames("id")
		c.SetParamValues("1")
		err := controller.Delete(c)

		// Check the outbound response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		resName, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.name")
		assert.Equal(t, "testing name", resName)

		resEmail, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.email")
		assert.Equal(t, "testing@email.id", resEmail)

		resPassword, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.password")
		assert.Equal(t, "Password123!", resPassword)

		deletedAt, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.deleted_at")
		assert.NotNil(t, deletedAt)

	})

	t.Run("Delete_failed", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		errorMessage := errors.New("delete data failed")
		errEx := exceptions.NewBusinessLogicError(exceptions.EBL10005, errorMessage)

		userService.On("Delete", req.Context(), uint64(1)).Return(nil, errEx)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.Delete(c)

		assert.Error(t, err)

		errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(err)
		assert.Equal(t, http.StatusUnprocessableEntity, errHttpCode)
		assert.Equal(t, exceptions.EBL10005, errBusinessCode)
		assert.Equal(t, "delete data failed", errMsg)
		assert.Nil(t, errOut)

	})

}

func TestUserController_GetById(t *testing.T) {

	t.Run("Get_success", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		userServiceResult := ent.User{
			ID:       1,
			RoleID:   1,
			Name:     "testing name",
			Email:    "testing@email.id",
			Password: "Password123!",
			Avatar:   "xx/ava.png",
		}

		userService.On("GetById", req.Context(), uint64(1)).Return(&userServiceResult, nil)

		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.GetById(c)

		// Check the outbound response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		resName, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.name")
		assert.Equal(t, "testing name", resName)

		resEmail, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.email")
		assert.Equal(t, "testing@email.id", resEmail)

		resPassword, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.password")
		assert.Equal(t, "Password123!", resPassword)

	})

	t.Run("Get_failed", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		req := httptest.NewRequest(http.MethodGet, "/users/1", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		errorMessage := errors.New("get data failed")
		errEx := exceptions.NewBusinessLogicError(exceptions.EBL10006, errorMessage)

		userService.On("GetById", req.Context(), uint64(1)).Return(nil, errEx)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := controller.GetById(c)

		assert.Error(t, err)

		errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(err)
		assert.Equal(t, http.StatusUnprocessableEntity, errHttpCode)
		assert.Equal(t, exceptions.EBL10006, errBusinessCode)
		assert.Equal(t, "get data failed", errMsg)
		assert.Nil(t, errOut)

	})
}

func TestUserController_GetAll(t *testing.T) {

	t.Run("Get_All_success", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		// CreateTx a request.
		req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		userServiceResult1 := ent.User{
			ID:       1,
			RoleID:   1,
			Name:     "testing name 1",
			Email:    "testing1@email.id",
			Password: "Password123!",
			Avatar:   "xx/ava1.png",
		}

		userServiceResult2 := ent.User{
			ID:       1,
			RoleID:   1,
			Name:     "testing name 2",
			Email:    "testing2@email.id",
			Password: "Password123!",
			Avatar:   "xx/ava2.png",
		}

		mockListUser := make([]*ent.User, 0)
		mockListUser = append(mockListUser, &userServiceResult1, &userServiceResult2)

		userService.On("GetAll", req.Context()).Return(mockListUser, nil)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := controller.GetAll(c)

		// Check the outbound response:
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		resName1, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.0.name")
		assert.Equal(t, "testing name 1", resName1)

		resName2, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.1.name")
		assert.Equal(t, "testing name 2", resName2)

		resEmail1, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.0.email")
		assert.Equal(t, "testing1@email.id", resEmail1)

		resEmail2, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.1.email")
		assert.Equal(t, "testing2@email.id", resEmail2)

		resPassword1, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.0.password")
		assert.Equal(t, "Password123!", resPassword1)

		resPassword2, _ := apputils.GetFieldBytes(res.Body.Bytes(), "data.1.password")
		assert.Equal(t, "Password123!", resPassword2)

	})

	t.Run("Get_All_failed", func(t *testing.T) {

		userService := new(mock_service.UserService)
		controller := NewUserController(userService)
		e := test.InitEchoTest(t)

		// CreateTx a request.
		req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")

		errorMessage := errors.New("get data failed")
		errEx := exceptions.NewBusinessLogicError(exceptions.EBL10006, errorMessage)

		userService.On("GetAll", req.Context()).Return(nil, errEx)

		// Call the CreateTx method of the UserController.
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := controller.GetAll(c)

		assert.Error(t, err)

		errHttpCode, errBusinessCode, errMsg, errOut := validator.MapperErrorCode(err)
		assert.Equal(t, http.StatusUnprocessableEntity, errHttpCode)
		assert.Equal(t, exceptions.EBL10006, errBusinessCode)
		assert.Equal(t, "get data failed", errMsg)
		assert.Nil(t, errOut)

	})
}
