package repository

import (
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/mocks/test_helper"
	"testing"
	"time"
)

func TestRoleUserRepositoryImpl_CreateTx(t *testing.T) {

	clientTx, _, ctx := test_helper.TestDbConnectionTx(t)
	userRoleRepo := NewRoleUserRepositoryImpl()

	dummyRoleUser := ent.RoleUser{
		ID:        1,
		UserID:    1,
		RoleID:    1,
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := userRoleRepo.CreateTx(ctx, clientTx, dummyRoleUser)
	assert.NoError(t, err)
	assert.Equal(t, dummyRoleUser.ID, result.ID)
	assert.Equal(t, dummyRoleUser.UserID, result.UserID)
	assert.Equal(t, dummyRoleUser.RoleID, result.RoleID)
	assert.Equal(t, dummyRoleUser.CreatedBy, result.CreatedBy)

}

func TestRoleUserRepositoryImpl_Update(t *testing.T) {

	clientTx, _, ctx := test_helper.TestDbConnectionTx(t)
	userRoleRepo := NewRoleUserRepositoryImpl()

	dummyRoleUser := ent.RoleUser{
		UserID:    1,
		RoleID:    1,
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := userRoleRepo.CreateTx(ctx, clientTx, dummyRoleUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRoleUser.UserID, result.UserID)
	assert.Equal(t, dummyRoleUser.RoleID, result.RoleID)
	assert.Equal(t, dummyRoleUser.CreatedBy, result.CreatedBy)

	dummyUpdatedRoleUser := ent.RoleUser{
		ID:        1,
		UserID:    2,
		RoleID:    2,
		CreatedBy: "user",
		CreatedAt: time.Time{},
		UpdatedBy: "user",
		UpdatedAt: time.Time{},
	}

	resultUpdated, err := userRoleRepo.Update(ctx, clientTx, dummyUpdatedRoleUser, dummyRoleUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, dummyUpdatedRoleUser.ID, resultUpdated.ID)
	assert.Equal(t, dummyUpdatedRoleUser.UserID, resultUpdated.UserID)
	assert.Equal(t, dummyUpdatedRoleUser.RoleID, resultUpdated.RoleID)
	assert.Equal(t, dummyUpdatedRoleUser.CreatedBy, resultUpdated.CreatedBy)

}
