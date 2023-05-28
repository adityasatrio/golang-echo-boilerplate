package service

import (
	"context"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"myapp/ent"
	"myapp/internal/applications/user/dto"
	mock_repository "myapp/mocks/role/repository"
	mock_repository3 "myapp/mocks/role_user/repository"
	mock_transaction "myapp/mocks/transaction"
	mock_repository2 "myapp/mocks/user/repository"
	"myapp/test/test_helper"
	"testing"
	"time"
)

var mockUserRepository = new(mock_repository2.UserRepository)
var mockRoleRepository = new(mock_repository.RoleRepository)
var mockRoleUserRepository = new(mock_repository3.RoleUserRepository)
var mockTransaction = new(mock_transaction.TrxService)

var service = NewUserServiceImpl(mockUserRepository, mockRoleRepository, mockRoleUserRepository, mockTransaction)

func getUserMock(id uint64, name string, email string, password string) ent.User {
	return ent.User{
		ID:    id,
		Name:  name,
		Email: email,
		//WARNING : becarefull with bool value, it always have default value as FALSE,
		//make sure when do testing DTO / request and actual mock or return value have same value

		// uncomment IsVerified will impact on failed test, because expected value false from default value,
		//if must true then need to adjust logic on service to alway set as true or get from method parameter / request
		//IsVerified: true,

		EmailVerifiedAt:  time.Time{},
		Password:         password,
		RememberToken:    "",
		SocialMediaID:    "",
		Avatar:           "",
		RoleID:           0,
		LoginType:        "",
		SubSpecialist:    "",
		FirebaseToken:    "",
		Info:             "",
		Description:      "",
		Specialist:       "",
		Phone:            "",
		LastAccessAt:     time.Time{},
		PregnancyMode:    false,
		LatestSkipUpdate: time.Time{},
		LatestDeletedAt:  time.Time{},
	}
}

func TestUserServiceImpl_Create(t *testing.T) {
	request := dto.UserRequest{
		RoleId:   0,
		Name:     "Admin",
		Email:    "admin@email.com",
		Password: "12345",
	}

	userMocks := []struct {
		name         string
		request      dto.UserRequest
		roleRequest  ent.RoleUser
		userRequest  ent.User
		userResponse ent.User
		scenario     bool
	}{
		{
			request:      request,
			roleRequest:  ent.RoleUser{UserID: 123000},
			name:         "Create_User_Success-1",
			userRequest:  getUserMock(uint64(0), "Admin", "admin@email.com", "12345"),
			userResponse: getUserMock(uint64(123000), "Admin", "admin@email.com", "12345"),
			scenario:     true,
		},
		{
			request:      request,
			roleRequest:  ent.RoleUser{UserID: 123001},
			name:         "Create_User_Success-2",
			userRequest:  getUserMock(uint64(0), "Admin", "admin@email.com", "12345"),
			userResponse: getUserMock(uint64(123001), "Admin", "admin@email.com", "12345"),
			scenario:     true,
		},
		{
			request:      request,
			roleRequest:  ent.RoleUser{UserID: 123001},
			name:         "Create_User_Failed-1",
			userRequest:  getUserMock(uint64(0), "Admin", "admin@email.com", "12345"),
			userResponse: getUserMock(uint64(123001), "Admin", "admin@email.com", "12345"),
			scenario:     false,
		},
	}

	_, txClient, ctx := test_helper.TestDbConnectionTx(t)

	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {

			if userMock.scenario {
				// Set up mock behavior
				mockTransaction.On("WithTx", ctx, mock.Anything).
					Run(func(args mock.Arguments) {

						fnTx := args.Get(1).(func(tx *ent.Tx) error)

						errTx := fnTx(txClient)
						require.NoError(t, errTx)
						require.NotNil(t, txClient.Client())

						if errTx != nil {
							return
						}

					}).
					Return(nil).
					Once()

				//the key for successful transaction mock is make sure `txClient` from withTx inner function use current struct
				mockUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.userRequest).
					Return(&userMock.userResponse, nil)

				mockRoleUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.roleRequest).
					Return(&userMock.roleRequest, nil)

				result, err := service.Create(ctx, &userMock.request)

				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, &userMock.userResponse, result)

			} else {
				mockTransaction.On("WithTx", ctx, mock.Anything).
					Run(func(args mock.Arguments) {
						fnTx := args.Get(1).(func(tx *ent.Tx) error)

						errTx := fnTx(txClient)
						require.NotNil(t, txClient.Client())
						if errTx != nil {
							return
						}
					}).
					Return(errors.New("fake failed saved")). //this return is the key for `withTx` do rollback process
					Once()

				mockUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.userRequest).
					Return(&userMock.userResponse, nil)

				mockRoleUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.roleRequest).
					Panic("failed saved")

				result, err := service.Create(ctx, &request)
				assert.NotNil(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserServiceImpl_Update(t *testing.T) {
	request := dto.UserRequest{
		RoleId:   0,
		Name:     "User",
		Email:    "user@email.com",
		Password: "12345",
	}

	id := uint64(123000)
	userRequest := getUserMock(uint64(0), "User", "user@email.com", "12345")
	userResponse := getUserMock(uint64(123000), "User", "user@email.com", "12345")
	roleRequest := ent.RoleUser{UserID: 123000}

	_, txClient, ctx := test_helper.TestDbConnectionTx(t)

	t.Run("Update_User_Success", func(t *testing.T) {

		mockTransaction.On("WithTx", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				fnTx := args.Get(1).(func(tx *ent.Tx) error)

				errTx := fnTx(txClient)
				require.NoError(t, errTx)
				require.NotNil(t, txClient.Client())
				if errTx != nil {
					return
				}

			}).Return(nil).
			Once()

		mockUserRepository.On("UpdateTx", ctx, txClient.Client(), userRequest, id).
			Return(&userResponse, nil)

		mockRoleUserRepository.On("Update", ctx, txClient.Client(), roleRequest, id).
			Return(&roleRequest, nil)

		result, err := service.Update(ctx, id, &request)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &userResponse, result)
	})

	t.Run("Update_User_Failed_User", func(t *testing.T) {
		err := errors.New("failed saved user")

		mockTransaction.On("WithTx", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				fnTx := args.Get(1).(func(tx *ent.Tx) error)

				errTx := fnTx(txClient)
				require.NoError(t, errTx)
				require.NotNil(t, txClient.Client())
				if errTx != nil {
					return
				}
			}).Return(err).
			Once()

		mockUserRepository.On("UpdateTx", ctx, txClient.Client(), userRequest, id).
			Return(nil, err)

		result, err := service.Update(ctx, id, &request)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("Update_User_Failed_Role_User", func(t *testing.T) {
		err := errors.New("failed saved role")

		mockTransaction.On("WithTx", ctx, mock.Anything).
			Run(func(args mock.Arguments) {
				fnTx := args.Get(1).(func(tx *ent.Tx) error)

				errTx := fnTx(txClient)
				require.NoError(t, errTx)
				require.NotNil(t, txClient.Client())
				if errTx != nil {
					return
				}
			}).Return(err).
			Once()

		mockUserRepository.On("UpdateTx", ctx, txClient.Client(), userRequest, id).
			Return(&userResponse, nil)

		mockRoleUserRepository.On("Update", ctx, txClient.Client(), roleRequest, id).
			Return(nil, err)

		result, err := service.Update(ctx, id, &request)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUserServiceImpl_Delete(t *testing.T) {
	ctx := context.Background()
	userMock := getUserMock(uint64(123000), "User-1", "user1@email.com", "12345")
	t.Run("Delete_success", func(t *testing.T) {
		mockUserRepository.On("SoftDelete", ctx, uint64(123000)).Return(&userMock, nil).Once()
		result, err := service.Delete(context.Background(), userMock.ID)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Delete_failed", func(t *testing.T) {
		errorMessage := errors.New("failed got user")
		mockUserRepository.On("SoftDelete", ctx, uint64(123000)).Return(nil, errorMessage).Once()
		result, err := service.Delete(context.Background(), userMock.ID)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestUserServiceImpl_GetById(t *testing.T) {
	ctx := context.Background()
	userMocks := []struct {
		id       uint64
		name     string
		expected ent.User
	}{
		{
			id:       uint64(10),
			name:     "GetById_User-1",
			expected: getUserMock(uint64(123000), "User-1", "user1@email.com", "12345"),
		},
		{
			id:       uint64(10),
			name:     "GetById_User-2",
			expected: getUserMock(uint64(11000), "User-2", "user2@email.com", "12345"),
		},
	}

	//table test:
	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {
			mockUserRepository.On("GetById", ctx, uint64(10)).Return(&userMock.expected, nil).Once()
			result, err := service.GetById(ctx, userMock.id)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
	}

	//subtest failed:
	t.Run("GetById_failed", func(t *testing.T) {
		errorMessage := errors.New("failed got user")
		mockUserRepository.On("GetById", ctx, uint64(10)).Return(nil, errorMessage).Once()
		result, err := service.GetById(ctx, uint64(10))
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUserServiceImpl_GetAll(t *testing.T) {
	ctx := context.Background()
	//subtest success:
	t.Run("GetAll_success", func(t *testing.T) {
		user := getUserMock(uint64(13000), "user example", "example@email.com", "12345")

		mockListUser := make([]*ent.User, 0)
		mockListUser = append(mockListUser, &user)

		mockUserRepository.On("GetAll", ctx).Return(mockListUser, nil).Once()
		result, err := service.GetAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	//subtest failed:
	t.Run("GetAll_failed", func(t *testing.T) {
		errorMessage := errors.New("failed get all user")
		mockUserRepository.On("GetAll", ctx).Return(nil, errorMessage).Once()
		result, err := service.GetAll(ctx)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}
