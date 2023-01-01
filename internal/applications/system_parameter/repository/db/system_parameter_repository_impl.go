package db

import (
	"context"
	"myapp/ent"
	"myapp/ent/system_parameter"
)

type SystemParameterRepositoryImpl struct {
	client *ent.Client
}

func NewSystemParameterRepository(dbConn *ent.Client) *SystemParameterRepositoryImpl {
	return &SystemParameterRepositoryImpl{
		client: dbConn,
	}
}

func (r *SystemParameterRepositoryImpl) Create(ctx context.Context, newData ent.System_parameter) (*ent.System_parameter, error) {
	saved, err := r.client.System_parameter.
		Create().
		SetKey(newData.Key).
		SetValue(newData.Value).
		SetCreatedBy("admin").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (r *SystemParameterRepositoryImpl) Update(ctx context.Context, id int, updateData ent.System_parameter) (*ent.System_parameter, error) {
	saved, err := r.client.System_parameter.
		UpdateOneID(id).
		SetKey(updateData.Key).
		SetValue(updateData.Value).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}
func (r *SystemParameterRepositoryImpl) Delete(ctx context.Context, id int) (*ent.System_parameter, error) {
	//TODO : soft delete
	err := r.client.System_parameter.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (r *SystemParameterRepositoryImpl) GetById(ctx context.Context, id int) (*ent.System_parameter, error) {
	data, err := r.client.System_parameter.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SystemParameterRepositoryImpl) GetAll(ctx context.Context) ([]*ent.System_parameter, error) {
	data, err := r.client.System_parameter.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SystemParameterRepositoryImpl) GetByKey(ctx context.Context, key string) (*ent.System_parameter, error) {
	data, err := r.client.System_parameter.Query().
		Where(system_parameter.KeyEqualFold(key)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

//TODO : add base schema
/*func (repository *SystemParameterRepositoryImpl) GetByIdNotDeleted(ctx context.Context, id int) (*ent.System_parameter, error) {
	data, err := repository.client.System_parameter.Get(ctx, id)
	if err != nil {
		return nil, err
	}


	return data, nil
}

func (repository *SystemParameterRepositoryImpl) GetAllNotDeleted(ctx context.Context) ([]*ent.System_parameter, error) {
	data, err := repository.client.System_parameter.Query().All(ctx)
	if err != nil {
		return nil, err
	}


	return data, nil
}*/
