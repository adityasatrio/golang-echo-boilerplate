package repository

import (
	"context"
	"fmt"
	"myapp/ent"
)

type (
	SystemParameterRepositoryImpl struct {
		client *ent.Client
	}
)

func NewSystemParameterRepository(dbConn *ent.Client) SystemParameterRepository {
	return &SystemParameterRepositoryImpl{
		client: dbConn,
	}
}

func (repo *SystemParameterRepositoryImpl) Create(ctx context.Context) (*ent.System_parameter, error) {
	saved, err := repo.client.System_parameter.
		Create().
		SetKey("aaa").
		SetValue("aaa").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Print("repo error", err)
	return saved, nil
}
