package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	roleRepository "myapp/internal/applications/role/repository"
	roleUserRepository "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/dto"
	userRepository "myapp/internal/applications/user/repository"
	"time"
)

type UserServiceImpl struct {
	userRepository     userRepository.UserRepository
	roleRepository     roleRepository.RoleRepository
	roleUserRepository roleUserRepository.RoleUserRepository
	transaction        transaction.TrxService
}

func NewUserServiceImpl(repository userRepository.UserRepository, roleRepository roleRepository.RoleRepository, roleUserRepository roleUserRepository.RoleUserRepository, transaction transaction.TrxService) *UserServiceImpl {
	return &UserServiceImpl{userRepository: repository, roleRepository: roleRepository, roleUserRepository: roleUserRepository, transaction: transaction}
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
			PregnancyMode: false,
		}

		//save user:

		userResult, err := s.userRepository.Create(ctx, tx, userRequest)
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
		_, errRoleUser := s.roleUserRepository.Create(ctx, tx, roleUserRequest)
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
		userResult, err := s.userRepository.Update(ctx, tx, userRequest, id)
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
		_, errRoleUser := s.roleUserRepository.Update(ctx, tx, roleUserRequest, id)
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
	data, err := s.userRepository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return data, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := s.userRepository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return data, nil
}

func (s *UserServiceImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	result, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	result, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}
