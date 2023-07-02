package repository

import (
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/test"
	"testing"
	"time"
)

func TestRoleUserRepositoryImpl_GetByUserIdAndRoleId(t *testing.T) {

	clientTx, _, ctx := test.DbConnectionTx(t)
	userRoleRepo := NewRoleUserRepositoryImpl(clientTx)

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

	getResult, err := userRoleRepo.GetByUserIdAndRoleId(ctx, dummyRoleUser.UserID, dummyRoleUser.RoleID)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, getResult.ID)
	assert.Equal(t, result.UserID, getResult.UserID)
	assert.Equal(t, result.RoleID, getResult.RoleID)
	assert.Equal(t, result.CreatedBy, getResult.CreatedBy)

	t.Cleanup(func() {
		test.DbConnectionClose(clientTx)
	})

}

func TestRoleUserRepositoryImpl_CreateTx(t *testing.T) {

	clientTx, _, ctx := test.DbConnectionTx(t)
	userRoleRepo := NewRoleUserRepositoryImpl(clientTx)

	dummyRoleUser := ent.RoleUser{
		UserID:    2,
		RoleID:    2,
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := userRoleRepo.CreateTx(ctx, clientTx, dummyRoleUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRoleUser.UserID, result.UserID)
	assert.Equal(t, dummyRoleUser.RoleID, result.RoleID)
	assert.Equal(t, dummyRoleUser.CreatedBy, result.CreatedBy)

	t.Cleanup(func() {
		test.DbConnectionClose(clientTx)
	})

}

func TestRoleUserRepositoryImpl_Update(t *testing.T) {

	clientTx, _, ctx := test.DbConnectionTx(t)
	userRoleRepo := NewRoleUserRepositoryImpl(clientTx)

	dummyRoleUser := ent.RoleUser{
		UserID:    3,
		RoleID:    3,
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := userRoleRepo.CreateTx(ctx, clientTx, dummyRoleUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRoleUser.UserID, result.UserID)
	assert.Equal(t, dummyRoleUser.RoleID, result.RoleID)
	assert.Equal(t, dummyRoleUser.CreatedBy, result.CreatedBy)

	result.UserID = 4
	result.RoleID = 4

	resultUpdated, err := userRoleRepo.UpdateTx(ctx, clientTx, result)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, resultUpdated.ID)
	assert.Equal(t, uint64(4), resultUpdated.UserID)
	assert.Equal(t, uint64(4), resultUpdated.RoleID)

	t.Cleanup(func() {
		test.DbConnectionClose(clientTx)
	})

}
