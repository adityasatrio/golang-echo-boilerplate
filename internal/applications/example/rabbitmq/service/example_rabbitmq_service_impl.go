package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/applications/system_parameter/repository/db"
)

type ExampleRabbitMQServiceImpl struct {
	repository db.SystemParameterRepository
}

func NewExampleRabbitMQService(repository db.SystemParameterRepository) *ExampleRabbitMQServiceImpl {
	return &ExampleRabbitMQServiceImpl{repository: repository}
}

func (t *ExampleRabbitMQServiceImpl) GetMessage(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error) {

	newData := ent.SystemParameter{
		Key:   create.Key,
		Value: create.Value,
	}

	exist, err := t.repository.GetByKey(ctx, newData.Key)
	if exist != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataAlreadyExist, err)
	}

	result, err := t.repository.Create(ctx, &newData)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	return result, nil
}
