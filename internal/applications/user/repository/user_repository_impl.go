package repository

import (
	"context"
	"myapp/ent"
	"myapp/ent/user"
	"time"
)

type UserRepositoryImpl struct {
	client *ent.Client
}

func NewUserRepositoryImpl(client *ent.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{client: client}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, client *ent.Tx, request ent.User) (*ent.User, error) {
	response, err := client.User.Create().
		SetRoleID(request.RoleID).
		SetName(request.Name).
		SetEmail(request.Email).
		SetPassword(request.Password).
		SetIsVerified(request.IsVerified).
		SetAvatar(request.Avatar).
		SetLastAccessAt(time.Now()).
		SetPregnancyMode(request.PregnancyMode).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, client *ent.Tx, request ent.User, id uint64) (*ent.User, error) {
	saved, err := client.User.
		UpdateOneID(id).
		SetRoleID(request.RoleID).
		SetName(request.Name).
		SetEmail(request.Email).
		SetPassword(request.Password).
		SetIsVerified(request.IsVerified).
		SetAvatar(request.Avatar).
		SetLastAccessAt(request.LastAccessAt).
		SetPregnancyMode(request.PregnancyMode).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, client *ent.Tx, id uint64) (*ent.User, error) {
	err := client.Role.DeleteOneID(id).Exec(ctx)

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
	data, err := r.client.User.Query().
		Where(user.And(
			user.ID(id),
			user.DeletedAtIsNil(),
		)).
		Only(ctx)
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
