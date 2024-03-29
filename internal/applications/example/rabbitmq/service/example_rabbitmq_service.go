package service

import (
	"context"
	"myapp/ent"
	"myapp/internal/applications/system_parameter/dto"
)

type ExampleRabbitMQService interface {
	GetMessage(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error)
}
