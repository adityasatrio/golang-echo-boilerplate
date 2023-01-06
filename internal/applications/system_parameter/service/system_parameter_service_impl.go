package service

import (
	"context"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/helper"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
	"time"
)

type SystemParameterServiceImpl struct {
	repository db.SystemParameterRepository
	cache      *cache.ChainCache[any]
}

func NewSystemParameterService(repository db.SystemParameterRepository, cache *cache.ChainCache[any]) *SystemParameterServiceImpl {
	return &SystemParameterServiceImpl{
		repository: repository,
		cache:      cache,
	}
}

func (s *SystemParameterServiceImpl) Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error) {
	newData := ent.SystemParameter{
		Key:   create.Key,
		Value: create.Value,
	}

	exist, err := s.repository.GetByKey(ctx, newData.Key)
	if exist != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10001, err)
	}

	result, err := s.repository.Create(ctx, newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
	}

	key := helper.CacheKey(result, nil)
	//err = s.cache.Set(ctx, key, result, store.WithExpiration(5*time.Minute))
	//err = s.cache.Set(ctx, key, &caching.CacheValue{SystemParameter: result}, store.WithExpiration(5*time.Minute))
	err = s.cache.Set(ctx, key, &result, store.WithExpiration(5*time.Minute))
	log.Info("cache error", err)
	return result, nil
}

func (s *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error) {

	existId, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	existKey, err := s.repository.GetByKey(ctx, update.Key)
	if existKey != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10001, err)
	}

	newData := ent.SystemParameter{
		Key:   update.Key,
		Value: update.Value,
	}

	updated, err := s.repository.Update(ctx, existId.ID, newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10004, err)
	}

	return updated, nil
}

func (s *SystemParameterServiceImpl) Delete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	_, err = s.repository.Delete(ctx, exist.ID)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return exist, nil
}

func (s *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {

	key := helper.CacheKey(ent.SystemParameter{ID: id}, nil)
	existCached, _ := s.cache.Get(ctx, key)
	if existCached != nil {
		castedResult := existCached.(*ent.SystemParameter)
		return castedResult, nil
	}

	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {

	key := helper.CacheKey(ent.SystemParameter{ID: -1}, nil)
	existCached, _ := s.cache.Get(ctx, key)
	if existCached != nil {
		castedResult := existCached.([]*ent.SystemParameter)
		return castedResult, nil
	}

	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	_ = s.cache.Set(ctx, key, result, store.WithExpiration(5*time.Minute))
	return result, nil
}
