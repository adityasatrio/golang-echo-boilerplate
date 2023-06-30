package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/globalutils"
	"myapp/internal/applications/cache"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
	"time"
)

type SystemParameterServiceImpl struct {
	repository db.SystemParameterRepository
	cache      cache.CachingService
}

func NewSystemParameterService(repository db.SystemParameterRepository, cache cache.CachingService) *SystemParameterServiceImpl {
	return &SystemParameterServiceImpl{repository: repository, cache: cache}
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

	result, err := s.repository.Create(ctx, &newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
	}

	_, err = s.cache.Create(ctx, globalutils.CacheKeySysParamWithId(result.ID), result, time.Hour*3)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error) {

	existKey, err := s.repository.GetByKey(ctx, update.Key)
	if existKey != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10001, err)
	}

	existId, err := s.GetById(ctx, id)
	if err != nil || existId == nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	existId.Key = update.Key
	existId.Value = update.Value

	updated, err := s.repository.Update(ctx, existId)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10004, err)
	}

	_, err = s.cache.Create(ctx, globalutils.CacheKeySysParamWithId(updated.ID), updated, time.Hour*3)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10007, err)
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

	_, err = s.cache.Delete(ctx, globalutils.CacheKeySysParamWithId(id))
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10007, err)
	}

	return exist, nil
}

func (s *SystemParameterServiceImpl) SoftDelete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil || exist == nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	deleted, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	_, err = s.cache.Delete(ctx, globalutils.CacheKeySysParamWithId(id))
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10007, err)
	}

	return deleted, nil
}

func (s *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {

	systemParameterCache, err := s.cache.Get(ctx, globalutils.CacheKeySysParamWithId(id), &ent.SystemParameter{})
	if err != nil && err != redis.Nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10007, err)
	}

	if systemParameterCache != nil {
		return systemParameterCache.(*ent.SystemParameter), nil
	}

	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	_, err = s.cache.Create(ctx, globalutils.CacheKeySysParamWithId(id), result, time.Hour*3)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10007, err)
	}

	return result, nil

}

func (s *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {

	systemParameterCache, err := s.cache.Get(ctx, globalutils.CacheKeySysParams(), &[]*ent.SystemParameter{})
	if err != nil && err != redis.Nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	if systemParameterCache != nil {
		systemParameterResult := append([]*ent.SystemParameter(nil), *systemParameterCache.(*[]*ent.SystemParameter)...)
		return systemParameterResult, err
	}

	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	_, err = s.cache.Create(ctx, globalutils.CacheKeySysParams(), &result, time.Hour*3)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10007, err)
	}

	return result, nil
}
