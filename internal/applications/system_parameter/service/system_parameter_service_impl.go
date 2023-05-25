package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
)

type SystemParameterServiceImpl struct {
	repository db.SystemParameterRepository
}

func NewSystemParameterService(repository db.SystemParameterRepository) *SystemParameterServiceImpl {
	return &SystemParameterServiceImpl{
		repository: repository,
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

	result, err := s.repository.Create(ctx, &newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error) {

	existKey, err := s.repository.GetByKey(ctx, update.Key)
	if existKey != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10001, err)
	}

	existId, err := s.repository.GetById(ctx, id)
	if err != nil || existId == nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	existId.Key = update.Key
	existId.Value = update.Value

	updated, err := s.repository.Update(ctx, existId.ID, existId)
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

func (s *SystemParameterServiceImpl) SoftDelete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil || exist == nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	deleted, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return deleted, nil
}

func (s *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {
	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	return result, nil
}

func (s *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}
