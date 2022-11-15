package usecase

import (
	"context"
	"myapp/domains/system_parameter/entity"
)

type SystemParameterCase interface {
	Create(ctx context.Context) (*entity.SystemParameter, error)
	Update(ctx context.Context) (*entity.SystemParameter, error)
	Delete(ctx context.Context) error
	GetById(ctx context.Context) (*entity.SystemParameter, error)
	GetAll(ctx context.Context) ([]*entity.SystemParameter, error)
}

func Create(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func Update(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func Delete(ctx context.Context) error {
	return nil
}

func GetById(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func GetAll(ctx context.Context) ([]*entity.SystemParameter, error) {
	return nil, nil
}
