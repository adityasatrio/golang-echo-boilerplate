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
	"myapp/test"
	"testing"
	"time"
)

var mockUserRepository = new(mock_repository2.UserRepository)
var mockRoleRepository = new(mock_repository.RoleRepository)
var mockRoleUserRepository = new(mock_repository3.RoleUserRepository)
var mockTransaction = new(mock_transaction.TrxService)

var service = NewUserService(mockUserRepository, mockRoleRepository, mockRoleUserRepository, mockTransaction)

func getUserMock(id uint64, name string, email string, password string) ent.User {
	return ent.User{
		ID:      id,
		Version: int64(0),
		Name:    name,
		Email:   email,
		//WARNING : careful with bool value, it always has default value as FALSE,
		//make sure when do testing DTO / request and actual mock or return value have same value

		// uncomment IsVerified will impact on failed test, because expected value false from default value,
		//if you must true then need to adjust logic on service to always set as true or get from method parameter / request
		//IsVerified: true,

		EmailVerifiedAt:  time.Time{},
		Password:         password,
		RememberToken:    "",
		SocialMediaID:    "",
		Avatar:           "",
		RoleID:           uint64(0),
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
		//DeletedAt:        nil,
	}
}

func TestUserServiceImpl_Create_Success(t *testing.T) {
	request := dto.UserRequest{
		RoleId:   0,
		Name:     "Admin",
		Email:    "admin@email.com",
		Password: "12345",
	}

	userMocks := []struct {
		name                 string
		request              dto.UserRequest
		roleRequest          ent.RoleUser
		userServiceParameter ent.User
		userServiceReturn    ent.User
		scenario             bool
	}{
		{
			request:              request,
			roleRequest:          ent.RoleUser{UserID: 123000},
			name:                 "Create_User_Success-1",
			userServiceParameter: getUserMock(uint64(0), "Admin", "admin@email.com", "12345"),
			userServiceReturn:    getUserMock(uint64(123000), "Admin", "admin@email.com", "12345"),
			scenario:             true,
		},
		{
			request:              request,
			roleRequest:          ent.RoleUser{UserID: 123001},
			name:                 "Create_User_Success-2",
			userServiceParameter: getUserMock(uint64(0), "Admin", "admin@email.com", "12345"),
			userServiceReturn:    getUserMock(uint64(123001), "Admin", "admin@email.com", "12345"),
			scenario:             true,
		},
		{
			request:              request,
			roleRequest:          ent.RoleUser{UserID: 123001},
			name:                 "Create_User_Failed-1",
			userServiceParameter: getUserMock(uint64(0), "Admin", "admin@email.com", "12345"),
			userServiceReturn:    getUserMock(uint64(123001), "Admin", "admin@email.com", "12345"),
			scenario:             false,
		},
	}

	_, txClient, ctx := test.DbConnectionTx(t)

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
				mockUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.userServiceParameter).
					Return(&userMock.userServiceReturn, nil)

				mockRoleUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.roleRequest).
					Return(&userMock.roleRequest, nil)

				result, err := service.Create(ctx, &userMock.request)

				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, &userMock.userServiceReturn, result)

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

				mockUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.userServiceParameter).
					Return(&userMock.userServiceReturn, nil)

				mockRoleUserRepository.On("CreateTx", ctx, txClient.Client(), userMock.roleRequest).
					Panic("failed saved")

				result, err := service.Create(ctx, &request)
				assert.NotNil(t, err)
				assert.Nil(t, result)
			}
		})
	}

	defer func() {
		test.DbConnectionCloseTx(txClient)
	}()
}

func TestUserServiceImpl_Update_Success(t *testing.T) {
	requestUpdate := dto.UserRequest{
		RoleId:   0,
		Name:     "User update",
		Email:    "user_update@email.com",
		Password: "12345_update",
	}

	_, txClient, ctx := test.DbConnectionTx(t)

	id := uint64(123000)
	userExisting := getUserMock(uint64(123000), "User", "user@email.com", "12345")
	userUpdated := getUserMock(uint64(123000), "User update", "user_update@email.com", "12345_update")
	userRoleExisting := ent.RoleUser{UserID: 123000, RoleID: uint64(0)}
	userRoleUpdated := ent.RoleUser{UserID: 123000, RoleID: uint64(1)}

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

	mockUserRepository.On("GetById", ctx, id).Return(&userExisting, nil)
	mockRoleUserRepository.On("GetByUserIdAndRoleId", ctx, userExisting.ID, userExisting.RoleID).Return(&userRoleExisting, nil)
	mockUserRepository.On("UpdateTx", ctx, txClient.Client(), &userExisting).Return(&userUpdated, nil)

	userRoleExisting.UserID = userExisting.ID
	userRoleExisting.RoleID = uint64(1)
	mockRoleUserRepository.On("UpdateTx", ctx, txClient.Client(), &userRoleExisting).Return(&userRoleUpdated, nil)

	result, err := service.Update(ctx, id, &requestUpdate)
	assert.NoError(t, err)
	assert.Equal(t, requestUpdate.Name, result.Name)
	assert.Equal(t, requestUpdate.Email, result.Email)
	assert.Equal(t, requestUpdate.Password, result.Password)

	defer func() {
		test.DbConnectionCloseTx(txClient)
	}()
}

func TestUserServiceImpl_Update_UserFailed(t *testing.T) {
	requestUpdate := dto.UserRequest{
		RoleId:   0,
		Name:     "User2 update",
		Email:    "user2_update@email.com",
		Password: "12345_update",
	}

	_, txClient, ctx := test.DbConnectionTx(t)

	id := uint64(123002)
	userExisting := getUserMock(uint64(123002), "User2", "user2@email.com", "12345")
	userRoleExisting := ent.RoleUser{UserID: 123002, RoleID: uint64(0)}

	err := errors.New("failed saved user")
	mockTransaction.On("WithTx", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			fnTx := args.Get(1).(func(tx *ent.Tx) error)

			errTx := fnTx(txClient)
			require.Error(t, errTx)
			require.NotNil(t, txClient.Client())
			if errTx != nil {
				return
			}
		}).Return(err).
		Once()

	mockUserRepository.On("GetById", ctx, id).Return(&userExisting, nil)
	mockRoleUserRepository.On("GetByUserIdAndRoleId", ctx, userExisting.ID, userExisting.RoleID).Return(&userRoleExisting, nil)
	mockUserRepository.On("UpdateTx", ctx, txClient.Client(), &userExisting).Return(nil, err)

	result, err := service.Update(ctx, id, &requestUpdate)
	assert.Error(t, err)
	assert.Nil(t, result)

	defer func() {
		test.DbConnectionCloseTx(txClient)
	}()

}

func TestUserServiceImpl_Update_UserRoleFailed(t *testing.T) {
	requestUpdate := dto.UserRequest{
		RoleId:   0,
		Name:     "User3 update",
		Email:    "user3_update@email.com",
		Password: "12345_update",
	}

	_, txClient, ctx := test.DbConnectionTx(t)

	id := uint64(123003)
	userExisting := getUserMock(id, "User3", "user3@email.com", "12345")
	userRoleExisting := ent.RoleUser{UserID: id, RoleID: uint64(0)}

	err := errors.New("failed saved role")

	mockTransaction.On("WithTx", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			fnTx := args.Get(1).(func(tx *ent.Tx) error)

			errTx := fnTx(txClient)
			require.Error(t, errTx)
			require.NotNil(t, txClient.Client())
			if errTx != nil {
				return
			}
		}).Return(err).
		Once()

	mockUserRepository.On("GetById", ctx, id).Return(&userExisting, nil)
	mockRoleUserRepository.On("GetByUserIdAndRoleId", ctx, userExisting.ID, userExisting.RoleID).Return(&userRoleExisting, nil)
	mockUserRepository.On("UpdateTx", ctx, txClient.Client(), &userExisting).Return(nil, err)

	userRoleExisting.UserID = userExisting.ID
	userRoleExisting.RoleID = uint64(1)
	mockRoleUserRepository.On("UpdateTx", ctx, txClient.Client(), &userRoleExisting, id).Return(nil, err)

	result, err := service.Update(ctx, id, &requestUpdate)
	assert.Error(t, err)
	assert.Nil(t, result)

	defer func() {
		test.DbConnectionCloseTx(txClient)
	}()
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
