package db

import (
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/test/test_helper"
	"testing"
)

func TestSystemParameterRepositoryImpl_Create(t *testing.T) {

	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key1",
		Value: "value1",
	}

	result, err := repo.Create(ctx, &createUser)

	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})

}

func TestSystemParameterRepositoryImpl_Update(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key10001",
		Value: "value10001",
	}

	result, err := repo.Create(ctx, &createUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	resultId, err := repo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, resultId.Key)
	assert.Equal(t, createUser.Value, resultId.Value)

	resultId.Key = "key20002"
	resultId.Value = "value20002"

	affected, err := repo.Update(ctx, resultId)
	assert.NoError(t, err)
	assert.Equal(t, 1, affected)

	afterUpdated, err := repo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, afterUpdated.ID)
	assert.Equal(t, "key20002", afterUpdated.Key)
	assert.Equal(t, "value20002", afterUpdated.Value)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})

}

func TestSystemParameterRepositoryImpl_Delete(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key000123",
		Value: "value000123",
	}

	result, err := repo.Create(ctx, &createUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	deletedResult, err := repo.Delete(ctx, result.ID)
	assert.NoError(t, err)
	assert.Nil(t, deletedResult)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestSystemParameterRepositoryImpl_SoftDelete(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key1XXX1",
		Value: "value1XXX1",
	}

	result, err := repo.Create(ctx, &createUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	deletedResult, err := repo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deletedResult.DeletedAt)
	assert.NotNil(t, deletedResult.DeletedBy)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestSystemParameterRepositoryImpl_GetById(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key1YYY1",
		Value: "value1YYY1",
	}

	result, err := repo.Create(ctx, &createUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	getResult, err := repo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.Equal(t, result.ID, getResult.ID)
	assert.Equal(t, result.Key, getResult.Key)
	assert.Equal(t, result.Value, getResult.Value)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}

func TestSystemParameterRepositoryImpl_GetAll(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key1AAA1",
		Value: "value1AAA1",
	}

	result, err := repo.Create(ctx, &createUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	getResultAll, err := repo.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, getResultAll)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})

}

func TestSystemParameterRepositoryImpl_GetByKey(t *testing.T) {
	client, ctx := test_helper.TestDbConnection(t)

	repo := NewSystemParameterRepository(client)

	createUser := ent.SystemParameter{
		Key:   "key1UNIQUE1",
		Value: "value1UNIQUE1",
	}

	result, err := repo.Create(ctx, &createUser)
	assert.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, createUser.Key, result.Key)
	assert.Equal(t, createUser.Value, result.Value)

	getResult, err := repo.GetByKey(ctx, "key1UNIQUE1")
	assert.NoError(t, err)
	assert.Equal(t, result.ID, getResult.ID)
	assert.Equal(t, result.Key, getResult.Key)
	assert.Equal(t, result.Value, getResult.Value)

	t.Cleanup(func() {
		test_helper.TestDbConnectionClose(client)
	})
}
