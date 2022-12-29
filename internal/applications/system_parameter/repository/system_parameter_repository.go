package repository

import (
	"context"
	"myapp/ent"
)

type SystemParameterRepository interface {
	Create(ctx context.Context, newData ent.System_parameter) (*ent.System_parameter, error)
	Update(ctx context.Context, id int, updateData ent.System_parameter) (*ent.System_parameter, error)
	Delete(ctx context.Context, id int) (*ent.System_parameter, error)
	GetById(ctx context.Context, id int) (*ent.System_parameter, error)
	GetAll(ctx context.Context) ([]*ent.System_parameter, error)
	//GetByIdAndNotDeleted(ctx context.Context, id int) (*ent.System_parameter, error)
	//GetAllAndNotDeleted(ctx context.Context) ([]*ent.System_parameter, error)
}
