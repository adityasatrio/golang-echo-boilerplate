package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/exceptions/errcode"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository"
)

type (
	SystemParameterServiceImpl struct {
		repository repository.SystemParameterRepository
	}
)

func NewSystemParameterService(repository repository.SystemParameterRepository) SystemParameterService {
	return &SystemParameterServiceImpl{
		repository: repository,
	}
}

func (service *SystemParameterServiceImpl) Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.System_parameter, error) {
	newData := ent.System_parameter{
		Key:   create.Key,
		Value: create.Value,
	}

	exist, err := service.repository.GetByKey(ctx, newData.Key)
	if exist != nil {
		return nil, exceptions.NewBusinessLogicError(errcode.EBL10001, err)
	}

	result, err := service.repository.Create(ctx, newData)
	if err != nil {
		return nil, exceptions.NewDataCreateError(err)
	}

	return result, nil
}

func (service *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.System_parameter, error) {

	exist, err := service.repository.GetByKey(ctx, update.Key)
	if exist != nil {
		return nil, exceptions.NewBusinessLogicError(errcode.EBL10001, err)
	}

	exist, err = service.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewDataNotFoundError(err)
	}

	newData := ent.System_parameter{
		Key:   update.Key,
		Value: update.Value,
	}

	updated, err := service.repository.Update(ctx, exist.ID, newData)
	if err != nil {
		return nil, exceptions.NewDataUpdateError(err)
	}

	return updated, nil
}

func (service *SystemParameterServiceImpl) Delete(ctx context.Context, id int) (*ent.System_parameter, error) {
	exist, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewDataNotFoundError(err)
	}

	_, err = service.repository.Delete(ctx, exist.ID)
	if err != nil {
		return nil, exceptions.NewDeleteError(err)
	}

	return exist, nil
}

func (service *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.System_parameter, error) {
	result, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewDataNotFoundError(err)
	}

	return result, nil
}

func (service *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.System_parameter, error) {
	result, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewDataGetError(err)
	}

	return result, nil
}
