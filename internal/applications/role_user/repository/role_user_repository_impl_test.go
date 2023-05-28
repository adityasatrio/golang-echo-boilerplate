package repository

import (
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/test/test_helper"
	"testing"
	"time"
)

func TestRoleUserRepositoryImpl_CreateTx(t *testing.T) {

	clientTx, _, ctx := test_helper.TestDbConnectionTx(t)
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

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(clientTx)
	})

}

func TestRoleUserRepositoryImpl_Update(t *testing.T) {

	clientTx, _, ctx := test_helper.TestDbConnectionTx(t)
	userRoleRepo := NewRoleUserRepositoryImpl(clientTx)

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
		RoleID:    1,
		CreatedBy: "user",
		CreatedAt: time.Time{},
		UpdatedBy: "user",
		UpdatedAt: time.Time{},
	}

	existingRoleUser, err := userRoleRepo.GetByUserIdAndRoleId(ctx, result.ID, result.RoleID)
	assert.NoError(t, err)

	existingRoleUser.

	resultUpdated, err := userRoleRepo.UpdateTx(ctx, clientTx, dummyUpdatedRoleUser)
	assert.NoError(t, err)
	assert.Equal(t, dummyUpdatedRoleUser.ID, resultUpdated.ID)
	assert.Equal(t, dummyUpdatedRoleUser.UserID, resultUpdated.UserID)
	assert.Equal(t, dummyUpdatedRoleUser.RoleID, resultUpdated.RoleID)
	assert.Equal(t, dummyUpdatedRoleUser.CreatedBy, resultUpdated.CreatedBy)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(clientTx)
	})

}
