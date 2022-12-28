package service

import (
	"context"
	"myapp/domains/system_parameter"
)

type (
	SystemParameterService interface {
		Hello(ctx context.Context) (string, error)

		Create(ctx context.Context) (*system_parameter.SystemParameter, error)
		Update(ctx context.Context) (*system_parameter.SystemParameter, error)
		Delete(ctx context.Context) error
		GetById(ctx context.Context) (*system_parameter.SystemParameter, error)
		GetAll(ctx context.Context) ([]*system_parameter.SystemParameter, error)
	}
)
