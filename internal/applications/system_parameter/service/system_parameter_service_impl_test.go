package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/vars"
	mock_db "myapp/mocks/applications/system_parameter/repository/db"
	mock_cache "myapp/mocks/component/cache"
	"testing"
)

func TestSystemParameterServiceImpl_Create(t *testing.T) {

	createRequest := dto.SystemParameterCreateRequest{
		Key:   "Key",
		Value: "Value",
	}

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.Cache)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()

	createSystemParameter := ent.SystemParameter{
		Key:   "Key",
		Value: "Value",
	}

	mockRepo.On("GetByKey", ctx, createSystemParameter.Key).
		Return(nil, nil)

	mockRepo.On("Create", ctx, &createSystemParameter).
		Return(&createSystemParameter, nil)

	mockCache.On("Create", ctx, CacheKeySysParamWithId(createSystemParameter.ID), &createSystemParameter, vars.GetTtlShortPeriod()).
		Return(true, nil)

	result, err := service.Create(ctx, &createRequest)

	assert.NoError(t, err)
	assert.Equal(t, createSystemParameter.Key, result.Key)
	assert.Equal(t, createSystemParameter.Value, result.Value)
}

func TestSystemParameterServiceImpl_Update(t *testing.T) {

	updateRequest := dto.SystemParameterUpdateRequest{
		Key:   "Key",
		Value: "Value",
	}

	createSystemParameter := ent.SystemParameter{
		ID:    1,
		Key:   "Key",
		Value: "Value",
	}

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.Cache)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()

	mockRepo.On("GetByKey", ctx, createSystemParameter.Key).Return(nil, nil)

	mockCache.On("Get", ctx, CacheKeySysParamWithId(createSystemParameter.ID), &ent.SystemParameter{}).
		Return(nil, nil)

	mockRepo.On("GetById", ctx, 1).Return(&createSystemParameter, nil)

	mockCache.On("Create", ctx, CacheKeySysParamWithId(createSystemParameter.ID), &createSystemParameter, vars.GetTtlShortPeriod()).
		Return(true, nil)

	mockRepo.On("Update", ctx, &createSystemParameter).Return(&createSystemParameter, nil)

	result, err := service.Update(ctx, 1, &updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, createSystemParameter.Key, result.Key)
	assert.Equal(t, createSystemParameter.Value, result.Value)
}

func TestSystemParameterServiceImpl_Delete(t *testing.T) {

	deleteSystemParameter := ent.SystemParameter{
		ID:    1,
		Key:   "Key",
		Value: "Value",
	}

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.Cache)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()
	mockRepo.On("GetById", ctx, deleteSystemParameter.ID).Return(&deleteSystemParameter, nil)
	mockRepo.On("Delete", ctx, deleteSystemParameter.ID).Return(nil, nil)
	mockCache.On("Delete", ctx, CacheKeySysParamWithId(deleteSystemParameter.ID)).
		Return(true, nil)

	result, err := service.Delete(ctx, deleteSystemParameter.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSystemParameterServiceImpl_SoftDelete(t *testing.T) {

	deleteSystemParameter := ent.SystemParameter{
		ID:    1,
		Key:   "Key",
		Value: "Value",
	}

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.Cache)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()
	mockRepo.On("GetById", ctx, deleteSystemParameter.ID).Return(&deleteSystemParameter, nil)
	mockRepo.On("SoftDelete", ctx, deleteSystemParameter.ID).Return(&deleteSystemParameter, nil)
	mockCache.On("Delete", ctx, CacheKeySysParamWithId(deleteSystemParameter.ID)).
		Return(true, nil)

	result, err := service.SoftDelete(ctx, deleteSystemParameter.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSystemParameterServiceImpl_GetById(t *testing.T) {

	getSystemParameter := ent.SystemParameter{
		ID:    1,
		Key:   "Key",
		Value: "Value",
	}

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.Cache)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()
	mockCache.On("Get", ctx, CacheKeySysParamWithId(getSystemParameter.ID), &ent.SystemParameter{}).
		Return(nil, nil)

	mockRepo.On("GetById", ctx, getSystemParameter.ID).Return(&getSystemParameter, nil)

	mockCache.On("Create", ctx, CacheKeySysParamWithId(getSystemParameter.ID), &getSystemParameter, vars.GetTtlShortPeriod()).
		Return(true, nil)

	result, err := service.GetById(ctx, getSystemParameter.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSystemParameterServiceImpl_GetAll(t *testing.T) {

	getAllSystemParameter := ent.SystemParameter{
		ID:    1,
		Key:   "Key",
		Value: "Value",
	}

	mockList := make([]*ent.SystemParameter, 0)
	mockList = append(mockList, &getAllSystemParameter)

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.Cache)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()

	mockCache.On("Get", ctx, CacheKeySysParams(), &[]*ent.SystemParameter{}).
		Return(nil, nil)
	mockRepo.On("GetAll", ctx).Return(mockList, nil)
	mockCache.On("Create", ctx, CacheKeySysParams(), &mockList, vars.GetTtlShortPeriod()).
		Return(true, nil)

	result, err := service.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
