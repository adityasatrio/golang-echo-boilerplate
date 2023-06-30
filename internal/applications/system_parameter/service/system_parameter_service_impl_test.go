package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/globalutils"
	"myapp/internal/applications/system_parameter/dto"
	mock_cache "myapp/mocks/cache"
	mock_db "myapp/mocks/system_parameter/repository/db"
	"testing"
	"time"
)

func TestSystemParameterServiceImpl_Create(t *testing.T) {

	createRequest := dto.SystemParameterCreateRequest{
		Key:   "Key",
		Value: "Value",
	}

	mockRepo := new(mock_db.SystemParameterRepository)
	mockCache := new(mock_cache.CachingService)
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

	mockCache.On("Create", ctx, globalutils.CacheKeySysParamWithId(createSystemParameter.ID), &createSystemParameter, time.Hour*3).
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
	mockCache := new(mock_cache.CachingService)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()

	mockRepo.On("GetByKey", ctx, createSystemParameter.Key).Return(nil, nil)

	mockCache.On("Get", ctx, globalutils.CacheKeySysParamWithId(createSystemParameter.ID), &ent.SystemParameter{}).
		Return(nil, nil)

	mockRepo.On("GetById", ctx, 1).Return(&createSystemParameter, nil)

	mockCache.On("Create", ctx, globalutils.CacheKeySysParamWithId(createSystemParameter.ID), &createSystemParameter, time.Hour*3).
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
	mockCache := new(mock_cache.CachingService)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()
	mockRepo.On("GetById", ctx, deleteSystemParameter.ID).Return(&deleteSystemParameter, nil)
	mockRepo.On("Delete", ctx, deleteSystemParameter.ID).Return(nil, nil)
	mockCache.On("Delete", ctx, globalutils.CacheKeySysParamWithId(deleteSystemParameter.ID)).
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
	mockCache := new(mock_cache.CachingService)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()
	mockRepo.On("GetById", ctx, deleteSystemParameter.ID).Return(&deleteSystemParameter, nil)
	mockRepo.On("SoftDelete", ctx, deleteSystemParameter.ID).Return(&deleteSystemParameter, nil)
	mockCache.On("Delete", ctx, globalutils.CacheKeySysParamWithId(deleteSystemParameter.ID)).
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
	mockCache := new(mock_cache.CachingService)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()
	mockCache.On("Get", ctx, globalutils.CacheKeySysParamWithId(getSystemParameter.ID), &ent.SystemParameter{}).
		Return(nil, nil)

	mockRepo.On("GetById", ctx, getSystemParameter.ID).Return(&getSystemParameter, nil)

	mockCache.On("Create", ctx, globalutils.CacheKeySysParamWithId(getSystemParameter.ID), &getSystemParameter, time.Hour*3).
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
	mockCache := new(mock_cache.CachingService)
	service := NewSystemParameterService(mockRepo, mockCache)

	ctx := context.Background()

	mockCache.On("Get", ctx, globalutils.CacheKeySysParams(), &[]*ent.SystemParameter{}).
		Return(nil, nil)
	mockRepo.On("GetAll", ctx).Return(mockList, nil)
	mockCache.On("Create", ctx, globalutils.CacheKeySysParams(), &mockList, time.Hour*3).
		Return(true, nil)

	result, err := service.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
