package service

import (
	"context"
	"myapp/ent"
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

	result, err := service.repository.Create(ctx, newData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *SystemParameterServiceImpl) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.System_parameter, error) {
	newData := ent.System_parameter{
		Key:   update.Key,
		Value: update.Value,
	}

	result, err := service.repository.Update(ctx, id, newData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *SystemParameterServiceImpl) Delete(ctx context.Context, id int) (*ent.System_parameter, error) {
	result, err := service.repository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *SystemParameterServiceImpl) GetById(ctx context.Context, id int) (*ent.System_parameter, error) {
	result, err := service.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*ent.System_parameter, error) {
	result, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
