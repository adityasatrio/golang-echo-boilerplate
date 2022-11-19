package usecase

import (
	"context"
	"myapp/domains/system_parameter/entity"
)

type (
	SystemParameterCase interface {
		Hello(ctx context.Context) (string, error)
		CreateSystemParameter(ctx context.Context) (*entity.SystemParameter, error)
		UpdateSystemParameter(ctx context.Context) (*entity.SystemParameter, error)
		DeleteSystemParameter(ctx context.Context) error
		GetSystemParameterById(ctx context.Context) (*entity.SystemParameter, error)
		GetSystemParameterAll(ctx context.Context) ([]*entity.SystemParameter, error)
	}

	SystemParameterCaseImpl struct {
		//inject repo
	}
)

// NewUseCase New constructor for DI
func NewUseCase( /*inject repo*/ ) *SystemParameterCaseImpl {
	return &SystemParameterCaseImpl{}
}

func (impl *SystemParameterCaseImpl) Hello(ctx context.Context) (string, error) {
	return "hello from case impl", nil
}

func (impl *SystemParameterCaseImpl) CreateSystemParameter(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func (impl *SystemParameterCaseImpl) UpdateSystemParameter(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func (impl *SystemParameterCaseImpl) DeleteSystemParameter(ctx context.Context) error {
	return nil
}

func (impl *SystemParameterCaseImpl) GetSystemParameterById(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func (impl *SystemParameterCaseImpl) GetSystemParameterAll(ctx context.Context) ([]*entity.SystemParameter, error) {
	return nil, nil
}
