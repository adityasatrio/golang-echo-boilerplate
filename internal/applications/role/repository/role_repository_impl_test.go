package repository

import (
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/test/test_helper"
	"testing"
	"time"
)

func TestRoleRepositoryImpl_Create(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)
	roleRepo := NewRoleRepositoryImpl(client)

	dummyRole := ent.Role{
		ID:        0,
		Name:      "john doe",
		Text:      "head heart wallets",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := roleRepo.Create(ctx, dummyRole)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRole.Name, result.Name)
	assert.Equal(t, dummyRole.Text, result.Text)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestRoleRepositoryImpl_Update(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)
	roleRepo := NewRoleRepositoryImpl(client)

	dummyRole := ent.Role{
		Name:      "john doe",
		Text:      "head heart wallets",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := roleRepo.Create(ctx, dummyRole)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRole.Name, result.Name)
	assert.Equal(t, dummyRole.Text, result.Text)

	result.Name = "john doe updated"
	result.Text = "head heart wallets updated"

	resultUpdated, err := roleRepo.Update(ctx, result)
	assert.NoError(t, err)
	assert.NotNil(t, resultUpdated.ID)
	assert.Equal(t, "john doe updated", resultUpdated.Name)
	assert.Equal(t, "head heart wallets updated", resultUpdated.Text)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestRoleRepositoryImpl_Delete(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)
	roleRepo := NewRoleRepositoryImpl(client)

	dummyRole := ent.Role{
		Name:      "john doe",
		Text:      "head heart wallets",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := roleRepo.Create(ctx, dummyRole)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRole.Name, result.Name)
	assert.Equal(t, dummyRole.Text, result.Text)

	resultUpdated, err := roleRepo.Delete(ctx, result.ID)
	assert.NoError(t, err)
	assert.Nil(t, resultUpdated)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestRoleRepositoryImpl_SoftDelete(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)
	roleRepo := NewRoleRepositoryImpl(client)

	dummyRole := ent.Role{
		Name:      "john doe",
		Text:      "head heart wallets",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := roleRepo.Create(ctx, dummyRole)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRole.Name, result.Name)
	assert.Equal(t, dummyRole.Text, result.Text)

	resultDeleted, err := roleRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultDeleted.ID)
	assert.NotNil(t, resultDeleted.DeletedBy)
	assert.NotNil(t, resultDeleted.DeletedAt)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestRoleRepositoryImpl_GetId(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)
	roleRepo := NewRoleRepositoryImpl(client)

	dummyRole := ent.Role{
		Name:      "john doe",
		Text:      "head heart wallets",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := roleRepo.Create(ctx, dummyRole)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRole.Name, result.Name)
	assert.Equal(t, dummyRole.Text, result.Text)

	resultGetId, err := roleRepo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultGetId)
	assert.Equal(t, resultGetId.ID, result.ID)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestRoleRepositoryImpl_GetAll(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)
	roleRepo := NewRoleRepositoryImpl(client)

	dummyRole := ent.Role{
		Name:      "john doe",
		Text:      "head heart wallets",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result, err := roleRepo.Create(ctx, dummyRole)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, dummyRole.Name, result.Name)
	assert.Equal(t, dummyRole.Text, result.Text)

	resultDeleted, err := roleRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultDeleted.ID)
	assert.NotNil(t, resultDeleted.DeletedBy)
	assert.NotNil(t, resultDeleted.DeletedAt)

	dummyRole2 := ent.Role{
		Name:      "john doe 2",
		Text:      "head heart wallets 2",
		CreatedBy: "user",
		CreatedAt: time.Time{},
	}

	result2, err := roleRepo.Create(ctx, dummyRole2)
	assert.NoError(t, err)
	assert.NotNil(t, result2.ID)
	assert.Equal(t, dummyRole2.Name, result2.Name)
	assert.Equal(t, dummyRole2.Text, result2.Text)

	resultAll, err := roleRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, result2.Name, resultAll[0].Name)
	assert.Equal(t, result2.Text, resultAll[0].Text)
	assert.NotNil(t, result2.ID)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})

}
