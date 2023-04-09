package db

import (
	"context"
	"myapp/ent"
)

type SystemParameterRepository interface {
	Create(ctx context.Context, newData *ent.SystemParameter) (*ent.SystemParameter, error)
	Update(ctx context.Context, id int, updateData *ent.SystemParameter) (*ent.SystemParameter, error)
	Delete(ctx context.Context, id int) (*ent.SystemParameter, error)
	GetById(ctx context.Context, id int) (*ent.SystemParameter, error)
	GetAll(ctx context.Context) ([]*ent.SystemParameter, error)
	GetByKey(ctx context.Context, key string) (*ent.SystemParameter, error)
	//GetByIdAndNotDeleted(ctx context.Context, id int) (*ent.System_parameter, error)
	//GetAllAndNotDeleted(ctx context.Context) ([]*ent.System_parameter, error)
}
