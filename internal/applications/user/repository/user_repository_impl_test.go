package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/mocks/test_helper"
	"testing"
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
				RoleID:     1,
				Name:       "John",
				Email:      "john@example.com",
				Password:   "password",
				Avatar:     "avatar",
				CreatedBy:  "user",
				IsVerified: false,
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

	client, ctx := test_helper.TestDbConnection(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewUserRepositoryImpl(client)
			result, err := repo.CreateTx(ctx, client, tt.newUser)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.newUser.Name, result.Name)
			assert.Equal(t, tt.newUser.Email, result.Email)

		})
	}

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		RoleID:    1,
		Password:  "password123",
		Avatar:    "avatar.jpg",
		CreatedBy: "user",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	// CreateTx an updated user with modified fields
	updatedUser := ent.User{
		RoleID:    2,
		Name:      "Jane Smith",
		Email:     "jane.smith@example.com",
		Password:  "newpassword",
		Avatar:    "new_avatar.jpg",
		CreatedBy: "user",
	}

	// UpdateTx the user using the repository method
	updated, err := userRepo.UpdateTx(ctx, client, updatedUser, result.ID)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, updated.Name, updatedUser.Name)
	assert.Equal(t, updated.Email, updatedUser.Email)
	assert.Equal(t, result.ID, updated.ID)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_Delete(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Name:      "John Doe delete",
		Email:     "john.doe.delete@example.com",
		RoleID:    1,
		Password:  "password123",
		Avatar:    "avatar.jpg",
		CreatedBy: "user",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.Delete(ctx, result.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_SoftDelete(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// CreateTx a test user
	createNewUserSoft := ent.User{
		Name:      "John Doe soft",
		Email:     "john.doe.soft@example.com",
		RoleID:    1,
		Password:  "password123",
		Avatar:    "avatar.jpg",
		CreatedBy: "user",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUserSoft)
	assert.NoError(t, err)
	assert.Equal(t, createNewUserSoft.Name, result.Name)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.NotNil(t, deleted.DeletedAt)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_GetId(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		RoleID:    1,
		Password:  "password123",
		Avatar:    "avatar.jpg",
		CreatedBy: "user",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	resultGetId, err := userRepo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultGetId)
	assert.Equal(t, resultGetId.ID, result.ID)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_GetAll(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepositoryImpl(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		RoleID:    1,
		Password:  "password123",
		Avatar:    "avatar.jpg",
		CreatedBy: "user",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Name, result.Name)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.NotNil(t, deleted.DeletedAt)

	// CreateTx a test user
	createNewUser2 := ent.User{
		Name:      "John Doe2",
		Email:     "john.doe2@example.com",
		RoleID:    1,
		Password:  "password123",
		Avatar:    "avatar.jpg",
		CreatedBy: "user",
	}

	result2, err := userRepo.CreateTx(ctx, client, createNewUser2)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser2.Name, result2.Name)
	assert.NotNil(t, result2.ID)

	resultAll, err := userRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser2.Name, resultAll[0].Name)
	assert.NotNil(t, result2.ID)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})

}
