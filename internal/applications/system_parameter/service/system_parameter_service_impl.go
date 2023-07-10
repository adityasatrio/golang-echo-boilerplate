package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/cache"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
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
		return nil, exceptions.NewBusinessLogicError(exceptions.DataAlreadyExist, err)
	}

	result, err := s.repository.Create(ctx, &newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	_, err = s.cache.Create(ctx, CacheKeySysParamWithId(result.ID), result, cache.CachingShortPeriod())
	if err != nil {
		//don't throw exception if create redis failed:
		return result, nil
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error) {

	_, err := s.repository.GetByKey(ctx, update.Key)

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

	_, err = s.cache.Create(ctx, CacheKeySysParamWithId(updated.ID), updated, cache.CachingShortPeriod())
	if err != nil {
		//don't throw exception if create redis failed:
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

	_, err = s.cache.Delete(ctx, CacheKeySysParamWithId(id))
	if err != nil {
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

	_, err = s.cache.Delete(ctx, CacheKeySysParamWithId(id))
	if err != nil {
		return deleted, nil
	}

	return deleted, nil
}

func (s *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {

	systemParameterCache, err := s.cache.Get(ctx, CacheKeySysParamWithId(id), &ent.SystemParameter{})
	if systemParameterCache != nil {
		return systemParameterCache.(*ent.SystemParameter), nil
	}

	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	_, err = s.cache.Create(ctx, CacheKeySysParamWithId(id), result, cache.CachingShortPeriod())
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	return result, nil

}

func (s *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {

	systemParameterCache, err := s.cache.Get(ctx, CacheKeySysParams(), &[]*ent.SystemParameter{})
	if systemParameterCache != nil {
		systemParameterResult := append([]*ent.SystemParameter(nil), *systemParameterCache.(*[]*ent.SystemParameter)...)
		return systemParameterResult, err
	}

	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	_, err = s.cache.Create(ctx, CacheKeySysParams(), &result, cache.CachingShortPeriod())
	if err != nil {
		return result, nil
	}

	return result, nil
}
