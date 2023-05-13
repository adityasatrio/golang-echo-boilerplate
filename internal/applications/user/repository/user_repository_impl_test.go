package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/mocks"
	"testing"
	"time"
)

func TestUserRepositoryImpl_Create(t *testing.T) {
	tests := []struct {
		name           string
		txClient       *ent.Client
		newUser        ent.User
		expectedResult *ent.User
		expectedError  error
	}{
		{
			name: "successful creation",
			newUser: ent.User{
				RoleID:        1,
				Name:          "John",
				Email:         "john@example.com",
				Password:      "password",
				IsVerified:    true,
				Avatar:        "avatar",
				PregnancyMode: false,
			},
			expectedResult: &ent.User{ID: 1},
			expectedError:  nil,
		},
		//{
		//	name:           "error creating user",
		//	newUser:        ent.User{},
		//	expectedResult: nil,
		//	expectedError:  errors.New("error creating user"),
		//},
	}

	client, ctx := mocks.TestConnection(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewUserRepositoryImpl(client)
			result, err := repo.Create(ctx, client, tt.newUser)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.newUser.Name, result.Name)
			assert.Equal(t, tt.newUser.Email, result.Email)

		})
	}

	defer func() {
		mocks.TestConnectionClose(client)
	}()
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	client, ctx := mocks.TestConnection(t)

	// Create a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// Create a test user
	createNewUser := ent.User{
		Name:          "John Doe",
		Email:         "john.doe@example.com",
		RoleID:        1,
		Password:      "password123",
		IsVerified:    true,
		Avatar:        "avatar.jpg",
		PregnancyMode: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := userRepo.Create(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	// Create an updated user with modified fields
	updatedUser := ent.User{
		RoleID:        2,
		Name:          "Jane Smith",
		Email:         "jane.smith@example.com",
		Password:      "newpassword",
		IsVerified:    false,
		Avatar:        "new_avatar.jpg",
		PregnancyMode: true,
	}

	// Update the user using the repository method
	updated, err := userRepo.Update(ctx, client, updatedUser, result.ID)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, updated.Name, updatedUser.Name)
	assert.Equal(t, updated.Email, updatedUser.Email)
	assert.Equal(t, result.ID, updated.ID)

	defer func() {
		mocks.TestConnectionClose(client)
	}()
}

func TestUserRepositoryImpl_Delete(t *testing.T) {

	client, tx, ctx := mocks.TestConnectionTx(t)

	// Create a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// Create a test user
	createNewUser := ent.User{
		Name:          "John Doe",
		Email:         "john.doe@example.com",
		RoleID:        1,
		Password:      "password123",
		IsVerified:    true,
		Avatar:        "avatar.jpg",
		PregnancyMode: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := userRepo.Create(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.Delete(ctx, tx, result.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)

	defer func() {
		mocks.TestConnectionClose(client)
	}()
}

func TestUserRepositoryImpl_SoftDelete(t *testing.T) {

	client, ctx := mocks.TestConnection(t)

	// Create a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// Create a test user
	createNewUserSoft := ent.User{
		Name:          "John Doe soft",
		Email:         "john.doe.soft@example.com",
		RoleID:        1,
		Password:      "password123",
		IsVerified:    true,
		Avatar:        "avatar.jpg",
		PregnancyMode: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := userRepo.Create(ctx, client, createNewUserSoft)
	assert.NoError(t, err)
	assert.Equal(t, createNewUserSoft.Name, result.Name)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.NotNil(t, deleted.DeletedAt)

	defer func() {
		mocks.TestConnectionClose(client)
	}()
}

func TestUserRepositoryImpl_GetId(t *testing.T) {

	client, ctx := mocks.TestConnection(t)

	// Create a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// Create a test user
	createNewUser := ent.User{
		Name:          "John Doe",
		Email:         "john.doe@example.com",
		RoleID:        1,
		Password:      "password123",
		IsVerified:    true,
		Avatar:        "avatar.jpg",
		PregnancyMode: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := userRepo.Create(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	resultId, err := userRepo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultId)
	assert.Equal(t, result.ID, resultId.ID)

	defer func() {
		mocks.TestConnectionClose(client)
	}()

}

func TestUserRepositoryImpl_GetAll(t *testing.T) {

	client, ctx := mocks.TestConnection(t)

	// Create a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// Create a test user
	createNewUser := ent.User{
		Name:          "John Doe",
		Email:         "john.doe@example.com",
		RoleID:        1,
		Password:      "password123",
		IsVerified:    true,
		Avatar:        "avatar.jpg",
		PregnancyMode: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := userRepo.Create(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.NotNil(t, deleted.DeletedAt)

	// Create a test user
	createNewUser2 := ent.User{
		Name:          "John Doe2",
		Email:         "john.doe2@example.com",
		RoleID:        1,
		Password:      "password123",
		IsVerified:    true,
		Avatar:        "avatar.jpg",
		PregnancyMode: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result2, err := userRepo.Create(ctx, client, createNewUser2)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser2.Name, result2.Name)
	assert.NotNil(t, result2.ID)

	resultAll, err := userRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser2.Name, resultAll[0].Name)
	assert.NotNil(t, result2.ID)

	defer func() {
		mocks.TestConnectionClose(client)
	}()

}
