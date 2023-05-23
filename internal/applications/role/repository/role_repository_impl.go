package repository

import (
	"context"
	"myapp/ent"
	"myapp/ent/role"
	"time"
)

type RoleRepositoryImpl struct {
	client *ent.Client
}

func NewRoleRepositoryImpl(client *ent.Client) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{client: client}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, role ent.Role) (*ent.Role, error) {

	response, err := r.client.Role.Create().
		SetName(role.Name).
		SetText(role.Text).
		SetCreatedBy(role.CreatedBy).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, role ent.Role, id uint64) (*ent.Role, error) {
	response, err := r.client.Role.
		UpdateOneID(id).
		SetName(role.Name).
		SetText(role.Text).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id uint64) (*ent.Role, error) {
	err := r.client.Role.DeleteOneID(id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *RoleRepositoryImpl) SoftDelete(ctx context.Context, id uint64) (*ent.Role, error) {
	deleted, err := r.client.Role.
		UpdateOneID(id).
		SetDeletedBy("user").
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return deleted, nil
}

func (r *RoleRepositoryImpl) GetById(ctx context.Context, id uint64) (*ent.Role, error) {
	data, err := r.client.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (r *RoleRepositoryImpl) GetAll(ctx context.Context) ([]*ent.Role, error) {
	data, err := r.client.Role.Query().
		Where(role.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
