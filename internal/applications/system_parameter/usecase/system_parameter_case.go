package usecase

import (
	"context"
	"myapp/domains/system_parameter/entity"
)

type (
	SystemParameterCase interface {
		Hello(ctx context.Context) (string, error)
		Create(ctx context.Context) (*entity.SystemParameter, error)
		Update(ctx context.Context) (*entity.SystemParameter, error)
		Delete(ctx context.Context) error
		GetById(ctx context.Context) (*entity.SystemParameter, error)
		GetAll(ctx context.Context) ([]*entity.SystemParameter, error)
	}

	systemParameterCase struct {
		//inject repo
	}
)

// NewUseCase New constructor for DI
func NewUseCase( /*inject repo*/ ) *systemParameterCase {
	return &systemParameterCase{}
}

func (impl *systemParameterCase) Hello(ctx context.Context) (string, error) {
	return "hello from case impl", nil
}

func (impl *systemParameterCase) Create(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func (impl *systemParameterCase) Update(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func (impl *systemParameterCase) Delete(ctx context.Context) error {
	return nil
}

func (impl *systemParameterCase) GetById(ctx context.Context) (*entity.SystemParameter, error) {
	return nil, nil
}

func (impl *systemParameterCase) GetAll(ctx context.Context) ([]*entity.SystemParameter, error) {
	return nil, nil
}
