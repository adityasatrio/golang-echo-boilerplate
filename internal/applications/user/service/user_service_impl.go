package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	repository2 "myapp/internal/applications/role/repository"
	repository3 "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/dto"
	"myapp/internal/applications/user/repository"
	"time"
)

type UserServiceImpl struct {
	repository         repository.UserRepository
	roleRepository     repository2.RoleRepository
	roleUserRepository repository3.RoleUserRepository

	transaction *transaction.TxService
}

func NewUserServiceImpl(repository repository.UserRepository, roleRepository repository2.RoleRepository,
	roleUserRepository repository3.RoleUserRepository,
	transaction *transaction.TxService) *UserServiceImpl {
	return &UserServiceImpl{repository: repository, roleRepository: roleRepository, roleUserRepository: roleUserRepository, transaction: transaction}
}

func (s *UserServiceImpl) Create(ctx context.Context, request *dto.UserRequest) (*ent.User, error) {

	//start transaction:
	var userNew = &ent.User{}
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {

		//create user object:
		userRequest := ent.User{
			RoleID:        request.RoleId,
			Name:          request.Name,
			Email:         request.Email,
			Password:      request.Password,
			IsVerified:    true,
			Avatar:        "",
			LastAccessAt:  time.Now(),
			PregnancyMode: false,
		}

		//save user:
		userResult, err := s.repository.Create(ctx, tx.Client(), userRequest)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
		}
		userNew = userResult

		//create role user object:
		roleUserRequest := ent.RoleUser{
			UserID: userNew.ID,
			RoleID: uint64(request.RoleId),
		}

		//save role_user:
		_, errRoleUser := s.roleUserRepository.Create(ctx, tx.Client(), roleUserRequest)
		if errRoleUser != nil {
			return exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return userNew, nil
}

func (s *UserServiceImpl) Update(ctx context.Context, id uint64, request *dto.UserRequest) (*ent.User, error) {

	//start transaction:
	var userUpdated = &ent.User{}
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {

		//create user object:
		userRequest := ent.User{
			RoleID:        request.RoleId,
			Name:          request.Name,
			Email:         request.Email,
			Password:      request.Password,
			IsVerified:    true,
			Avatar:        "",
			LastAccessAt:  time.Now(),
			PregnancyMode: false,
		}

		//update user:
		userResult, err := s.repository.Update(ctx, tx.Client(), userRequest, id)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
		}
		userUpdated = userResult

		//create role user object:
		roleUserRequest := ent.RoleUser{
			UserID: userUpdated.ID,
			RoleID: uint64(request.RoleId),
		}

		//update role_user:
		_, errRoleUser := s.roleUserRepository.Update(ctx, tx.Client(), roleUserRequest, id)
		if errRoleUser != nil {
			return exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return userUpdated, nil
}

func (s *UserServiceImpl) SoftDelete(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return data, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return data, nil
}

func (s *UserServiceImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}
