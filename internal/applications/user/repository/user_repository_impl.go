package repository

import (
	"context"
	"myapp/ent"
	"myapp/ent/user"
	"myapp/internal/applications/user/dto"
	"time"
)

type UserRepositoryImpl struct {
	client *ent.Client
}

func NewUserRepositoryImpl(client *ent.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{client: client}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, request dto.UserRequest) (*ent.User, error) {
	response, err := r.client.User.Create().
		SetRoleID(request.RoleId).
		SetName(request.Name).
		SetEmail(request.Email).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, request dto.UserRequest, id uint64) (*ent.User, error) {
	saved, err := r.client.User.
		UpdateOneID(id).
		SetRoleID(request.RoleId).
		SetName(request.Name).
		SetEmail(request.Email).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id uint64) (*ent.User, error) {
	err := r.client.Role.DeleteOneID(id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *UserRepositoryImpl) SoftDelete(ctx context.Context, id uint64) (*ent.User, error) {
	deleted, err := r.client.User.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return deleted, nil
}

func (r UserRepositoryImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (r UserRepositoryImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	data, err := r.client.User.Query().
		Where(user.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
