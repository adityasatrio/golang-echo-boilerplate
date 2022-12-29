package service

import (
	"context"
	"myapp/domains/system_parameter"
	"myapp/internal/applications/system_parameter/repository"
)

type (
	SystemParameterServiceImpl struct {
		//inject repository
		repository repository.SystemParameterRepository //interface
	}
)

// NewSystemParameterService New constructor for DI
func NewSystemParameterService(repository repository.SystemParameterRepository) SystemParameterService {
	return &SystemParameterServiceImpl{
		repository: repository,
	}
}

func (service *SystemParameterServiceImpl) Hello(ctx context.Context) (string, error) {
	/*result, err := service.repository.Create(ctx)
	if err != nil {
		fmt.Print("result ", result)
		fmt.Print("result err ", err)
	}
	*/
	return "hello from case impl", nil
}

func (service *SystemParameterServiceImpl) Create(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//result, err := service.repository.Create()

	return nil, nil
}

func (service *SystemParameterServiceImpl) Update(ctx context.Context) (*system_parameter.SystemParameter, error) {
	return nil, nil
}

func (service *SystemParameterServiceImpl) Delete(ctx context.Context) error {
	return nil
}

func (service *SystemParameterServiceImpl) GetById(ctx context.Context) (*system_parameter.SystemParameter, error) {
	return nil, nil
}

func (service *SystemParameterServiceImpl) GetAll(ctx context.Context) ([]*system_parameter.SystemParameter, error) {
	return nil, nil
}
