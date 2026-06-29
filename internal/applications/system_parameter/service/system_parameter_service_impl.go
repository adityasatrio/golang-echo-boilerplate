package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
	caching "myapp/internal/component/cache"
	"myapp/internal/vars"
)

type SystemParameterServiceImpl struct {
	repository db.SystemParameterRepository
	cache      caching.Cache
}

func NewSystemParameterService(repository db.SystemParameterRepository, cache caching.Cache) *SystemParameterServiceImpl {
	return &SystemParameterServiceImpl{repository: repository, cache: cache}
}

func (s *SystemParameterServiceImpl) Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error) {
	newData := ent.SystemParameter{
		Key:   create.Key,
		Value: create.Value,
	}

	exist, err := s.repository.GetByKey(ctx, newData.Key)
	if exist != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataAlreadyExist, err)
	}

	result, err := s.repository.Create(ctx, &newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	_, cacheErr := s.cache.Create(ctx, CacheKeySysParamWithId(result.ID), result, vars.GetTtlShortPeriod())
	if cacheErr != nil {
		return result, nil
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error) {

	_, _ = s.repository.GetByKey(ctx, update.Key)

	existId, err := s.GetById(ctx, id)
	if err != nil || existId == nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	existId.Key = update.Key
	existId.Value = update.Value

	updated, err := s.repository.Update(ctx, existId)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataUpdateFailed, err)
	}

	_, cacheErr := s.cache.Create(ctx, CacheKeySysParamWithId(updated.ID), updated, vars.GetTtlShortPeriod())
	if cacheErr != nil {
		return updated, nil
	}

	return updated, nil
}

func (s *SystemParameterServiceImpl) Delete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	_, err = s.repository.Delete(ctx, exist.ID)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataDeleteFailed, err)
	}

	_, cacheErr := s.cache.Delete(ctx, CacheKeySysParamWithId(id))
	if cacheErr != nil {
		return exist, nil
	}

	return exist, nil
}

func (s *SystemParameterServiceImpl) SoftDelete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil || exist == nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	deleted, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataDeleteFailed, err)
	}

	_, cacheErr := s.cache.Delete(ctx, CacheKeySysParamWithId(id))
	if cacheErr != nil {
		return deleted, nil
	}

	return deleted, nil
}

func (s *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {

	systemParameterCache, _ := s.cache.Get(ctx, CacheKeySysParamWithId(id), &ent.SystemParameter{})
	if sp, ok := systemParameterCache.(*ent.SystemParameter); ok && sp != nil {
		return sp, nil
	}

	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	_, err = s.cache.Create(ctx, CacheKeySysParamWithId(id), result, vars.GetTtlShortPeriod())
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	return result, nil

}

func (s *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {

	systemParameterCache, _ := s.cache.Get(ctx, CacheKeySysParams(), &[]*ent.SystemParameter{})
	if cached, ok := systemParameterCache.(*[]*ent.SystemParameter); ok && cached != nil {
		systemParameterResult := append([]*ent.SystemParameter(nil), *cached...)
		return systemParameterResult, nil
	}

	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	_, cacheErr := s.cache.Create(ctx, CacheKeySysParams(), &result, vars.GetTtlShortPeriod())
	if cacheErr != nil {
		return result, nil
	}

	return result, nil
}
