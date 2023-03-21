package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"myapp/ent"
	"myapp/mocks"
	"testing"
	"time"
)

var mockUserRepository = new(mocks.UserRepository)
var mockRoleRepository = new(mocks.RoleRepository)
var mockRoleUserRepository = new(mocks.RoleUserRepository)
var mockTransaction = new(mocks.TrxService)
var service = NewUserServiceImpl(mockUserRepository, mockRoleRepository, mockRoleUserRepository, mockTransaction)

func getUserMock(id uint64, name string, email string) ent.User {
	return ent.User{
		ID:               id,
		Name:             name,
		Email:            email,
		IsVerified:       false,
		EmailVerifiedAt:  time.Time{},
		Password:         "",
		RememberToken:    "",
		SocialMediaID:    "",
		Avatar:           "",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
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
		DeletedAt:        time.Time{},
		LatestSkipUpdate: time.Time{},
		LatestDeletedAt:  time.Time{},
	}
}

func TestUserServiceImpl_Create(t *testing.T) {
}

func TestUserServiceImpl_Update(t *testing.T) {
}

func TestUserServiceImpl_Delete(t *testing.T) {

}

func TestUserServiceImpl_GetById(t *testing.T) {

	userMocks := []struct {
		id       uint64
		name     string
		expected ent.User
	}{
		{
			id:       uint64(10),
			name:     "GetById_User-1",
			expected: getUserMock(uint64(123000), "User-1", "user1@email.com"),
		},
		{
			id:       uint64(10),
			name:     "GetById_User-2",
			expected: getUserMock(uint64(11000), "User-2", "user2@email.com"),
		},
	}

	//table test:
	for _, userMock := range userMocks {
		t.Run(userMock.name, func(t *testing.T) {
			mockUserRepository.On("GetById", mock.Anything, uint64(10)).Return(&userMock.expected, nil).Once()
			result, err := service.GetById(context.Background(), userMock.id)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
	}

	t.Run("GetById_failed", func(t *testing.T) {
		errorMessage := errors.New("failed got user")
		mockUserRepository.On("GetById", mock.Anything, uint64(10)).Return(nil, errorMessage).Once()
		result, err := service.GetById(context.Background(), uint64(10))
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestUserServiceImpl_GetAll(t *testing.T) {

	//sub test:
	t.Run("GetAll_success", func(t *testing.T) {
		user := getUserMock(uint64(13000), "user example", "example@email.com")

		mockListUser := make([]*ent.User, 0)
		mockListUser = append(mockListUser, &user)

		mockUserRepository.On("GetAll", mock.Anything).Return(mockListUser, nil).Once()
		result, err := service.GetAll(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetAll_failed", func(t *testing.T) {
		errorMessage := errors.New("failed get all user")
		mockUserRepository.On("GetAll", mock.Anything).Return(nil, errorMessage).Once()
		result, err := service.GetAll(context.Background())
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}
