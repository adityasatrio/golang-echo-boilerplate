package db

import (
	"context"
	"myapp/ent"
	"myapp/ent/systemparameter"
)

type SystemParameterRepositoryImpl struct {
	client *ent.Client
}

func NewSystemParameterRepository(dbConn *ent.Client) *SystemParameterRepositoryImpl {
	return &SystemParameterRepositoryImpl{
		client: dbConn,
	}
}

func (r *SystemParameterRepositoryImpl) Create(ctx context.Context, newData ent.SystemParameter) (*ent.SystemParameter, error) {
	saved, err := r.client.SystemParameter.
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

func (r *SystemParameterRepositoryImpl) Update(ctx context.Context, id int, updateData ent.SystemParameter) (*ent.SystemParameter, error) {
	saved, err := r.client.SystemParameter.
		UpdateOneID(id).
		SetKey(updateData.Key).
		SetValue(updateData.Value).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}
func (r *SystemParameterRepositoryImpl) Delete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	//TODO : soft delete
	err := r.client.SystemParameter.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (r *SystemParameterRepositoryImpl) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {
	data, err := r.client.SystemParameter.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SystemParameterRepositoryImpl) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {
	data, err := r.client.SystemParameter.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *SystemParameterRepositoryImpl) GetByKey(ctx context.Context, key string) (*ent.SystemParameter, error) {
	data, err := r.client.SystemParameter.Query().
		Where(systemparameter.KeyEqualFold(key)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}

//TODO : add base schema
/*func (repository *SystemParameterRepositoryImpl) GetByIdNotDeleted(ctx context.Context, id int) (*ent.SystemParameter, error) {
	data, err := repository.client.SystemParameter.Get(ctx, id)
	if err != nil {
		return nil, err
	}


	return data, nil
}

func (repository *SystemParameterRepositoryImpl) GetAllNotDeleted(ctx context.Context) ([]*ent.SystemParameter, error) {
	data, err := repository.client.SystemParameter.Query().All(ctx)
	if err != nil {
		return nil, err
	}


	return data, nil
}*/
