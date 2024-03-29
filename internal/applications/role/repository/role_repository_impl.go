package repository

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"myapp/ent"
	"myapp/ent/role"
	"time"
)

type RoleRepositoryImpl struct {
	client *ent.Client
}

func NewRoleRepository(client *ent.Client) *RoleRepositoryImpl {
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

func (r *RoleRepositoryImpl) Update(ctx context.Context, updateRole *ent.Role) (*ent.Role, error) {
	affected, err := r.client.Role.
		Update().Where(role.ID(updateRole.ID), role.Versions(updateRole.Versions)).
		SetName(updateRole.Name).
		SetText(updateRole.Text).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	if affected < 1 {
		log.Errorf("ID %d no records were updated in database", updateRole.ID)
		return nil, errors.New("no records were updated in database")
	}

	updatedRole, err := r.client.Role.Get(ctx, updateRole.ID)
	if err != nil {
		return nil, err
	}

	return updatedRole, nil
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
