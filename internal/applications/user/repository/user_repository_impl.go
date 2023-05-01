package repository

import (
	"context"
	"myapp/ent"
	"myapp/ent/user"
	"time"
)

type UserRepositoryImpl struct {
	client *ent.Client //non transactional client
}

func NewUserRepositoryImpl(client *ent.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{client: client}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, txClient *ent.Client, newUser ent.User) (*ent.User, error) {
	//txClient is transactional client that handled in service layer for post rollback logic
	response, err := txClient.User.Create().
		SetRoleID(newUser.RoleID).
		SetName(newUser.Name).
		SetEmail(newUser.Email).
		SetPassword(newUser.Password).
		SetIsVerified(newUser.IsVerified).
		SetAvatar(newUser.Avatar).
		SetLastAccessAt(time.Now()).
		SetPregnancyMode(newUser.PregnancyMode).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, txClient *ent.Client, updateUser ent.User, id uint64) (*ent.User, error) {
	saved, err := txClient.User.UpdateOneID(id).
		SetRoleID(updateUser.RoleID).
		SetName(updateUser.Name).
		SetEmail(updateUser.Email).
		SetPassword(updateUser.Password).
		SetIsVerified(updateUser.IsVerified).
		SetAvatar(updateUser.Avatar).
		SetLastAccessAt(time.Now()).
		SetPregnancyMode(updateUser.PregnancyMode).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *ent.Tx, id uint64) (*ent.User, error) {
	err := tx.Client().Role.DeleteOneID(id).Exec(ctx)

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

func (r *UserRepositoryImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {
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

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	data, err := r.client.User.Query().
		Where(user.DeletedAtIsNil()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil
}
