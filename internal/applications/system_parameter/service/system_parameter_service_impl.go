package service

import (
	"context"
	"github.com/eko/gocache/lib/v4/cache"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
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

func (s *SystemParameterServiceImpl) Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.System_parameter, error) {
	newData := ent.System_parameter{
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

	return result, nil
}

func (s *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.System_parameter, error) {

	existId, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	existKey, err := s.repository.GetByKey(ctx, update.Key)
	if existKey != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10001, err)
	}

	newData := ent.System_parameter{
		Key:   update.Key,
		Value: update.Value,
	}

	updated, err := s.repository.Update(ctx, existId.ID, newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10004, err)
	}

	return updated, nil
}

func (s *SystemParameterServiceImpl) Delete(ctx context.Context, id int) (*ent.System_parameter, error) {
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

func (s *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.System_parameter, error) {
	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.System_parameter, error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}
