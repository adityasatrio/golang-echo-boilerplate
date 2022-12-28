package repository

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/ent"
)

type (
	SystemParameterRepository interface {
		Create(ctx context.Context) (*ent.System_parameter, error)
	}

	repository struct {
		client *ent.Client
	}
)

func NewRepository( /*should inject DB init*/ ) *repository {
	return &repository{
		client: config.GetClient(),
	}
}

func (repo *repository) Create(ctx context.Context) (*ent.System_parameter, error) {
	saved, err := repo.client.System_parameter.
		Create().
		SetKey("abc").
		SetValue("abc").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Print("repo error", err)
	return saved, nil
}
