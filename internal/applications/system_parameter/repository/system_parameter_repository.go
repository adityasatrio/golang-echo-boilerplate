package repository

import (
	"context"
	"myapp/config"
	"myapp/ent"
)

type (
	SystemParameterRepository interface {
		create(ctx context.Context) (*ent.System_parameter, error)
	}

	repository struct {
		client *ent.Client
	}
)

func NewRepository() *repository {
	return &repository{
		client: config.GetClient(),
	}
}

func (repo *repository) create(ctx context.Context) (*ent.System_parameter, error) {
	saved, err := repo.client.System_parameter.
		Create().
		SetKey("abc").
		SetValue("abc").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}
