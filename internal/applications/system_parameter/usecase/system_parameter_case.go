package usecase

import (
	"context"
	"fmt"
	"myapp/domains/system_parameter/entity"
	"myapp/internal/applications/system_parameter/repository"
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
		repository repository.SystemParameterRepository //interface
	}
)

// NewUseCase New constructor for DI
func NewUseCase(repository repository.SystemParameterRepository) *systemParameterCase {
	return &systemParameterCase{
		repository: repository,
	}
}

func (impl *systemParameterCase) Hello(ctx context.Context) (string, error) {
	result, err := impl.repository.Create(ctx)
	if err != nil {
		fmt.Print("result ", result)
		fmt.Print("result err ", err)
	}

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
