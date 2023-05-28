package db

import (
	"context"
	"myapp/ent"
	"myapp/ent/systemparameter"
	"time"
)

type SystemParameterRepositoryImpl struct {
	client *ent.Client
}

func NewSystemParameterRepository(dbConn *ent.Client) *SystemParameterRepositoryImpl {
	return &SystemParameterRepositoryImpl{
		client: dbConn,
	}
}

func (r *SystemParameterRepositoryImpl) Create(ctx context.Context, newData *ent.SystemParameter) (*ent.SystemParameter, error) {
	saved, err := r.client.SystemParameter.
		Create().
		SetKey(newData.Key).
		SetValue(newData.Value).
		SetCreatedBy("user").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (r *SystemParameterRepositoryImpl) Update(ctx context.Context, updateData *ent.SystemParameter) (int, error) {
	affected, err := r.client.SystemParameter.
		Update().Where(systemparameter.ID(updateData.ID), systemparameter.Version(updateData.Version)).
		SetKey(updateData.Key).
		SetValue(updateData.Value).
		Save(ctx)

	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (r *SystemParameterRepositoryImpl) Delete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	err := r.client.SystemParameter.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *SystemParameterRepositoryImpl) SoftDelete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	updated, err := r.client.SystemParameter.
		UpdateOneID(id).
		SetDeletedBy("user").
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return updated, nil
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
