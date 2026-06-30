package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"myapp/ent"
	"myapp/exceptions"
	roleRepository "myapp/internal/applications/role/repository"
	roleUserRepository "myapp/internal/applications/role_user/repository"
	"myapp/internal/applications/user/dto"
	userRepository "myapp/internal/applications/user/repository"
	caching "myapp/internal/component/cache"
	"myapp/internal/component/transaction"
	"myapp/internal/vars"
)

type UserServiceImpl struct {
	userRepository     userRepository.UserRepository
	roleRepository     roleRepository.RoleRepository
	roleUserRepository roleUserRepository.RoleUserRepository
	transaction        transaction.Trx
	cache              caching.Cache
}

func NewUserService(userRepository userRepository.UserRepository, roleRepository roleRepository.RoleRepository, roleUserRepository roleUserRepository.RoleUserRepository, transaction transaction.Trx, cache caching.Cache) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository, roleRepository: roleRepository, roleUserRepository: roleUserRepository, transaction: transaction, cache: cache}
}

func (s *UserServiceImpl) Create(ctx context.Context, request *dto.UserRequest) (*ent.User, error) {

	// start transaction:
	var userNew = &ent.User{}
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
		}

		// create user object:
		userRequest := ent.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashedPassword),
			Avatar:   "",
			RoleID:   request.RoleId,
		}

		// save user:
		userResult, err := s.userRepository.CreateTx(ctx, tx.Client(), userRequest)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
		}
		userNew = userResult

		// create role user object:
		roleUserRequest := ent.RoleUser{
			UserID: userNew.ID,
			RoleID: request.RoleId,
		}

		// save role_user:
		_, errRoleUser := s.roleUserRepository.CreateTx(ctx, tx.Client(), roleUserRequest)
		if errRoleUser != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
		}

		// create cache, don't throw exception if failed:
		_, _ = s.cache.Create(ctx, CacheKeyUserWithId(userNew.ID), userNew, vars.GetTtlShortPeriod())

		return nil

	}); err != nil {
		// add rollback logic here
		log.Error("do rollback from transactional database operation")
		return nil, err
	}

	return userNew, nil
}

func (s *UserServiceImpl) Update(ctx context.Context, id uint64, request *dto.UserRequest) (*ent.User, error) {

	// start transaction:
	var userUpdated = &ent.User{}
	if err := s.transaction.WithTx(ctx, func(tx *ent.Tx) error {

		userExisting, err := s.userRepository.GetById(ctx, id)
		if userExisting == nil || err != nil {
			// log.Errorf("user data is not exist ID = %d", id)
			return exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
		}

		// get existing data for role_user based on old value
		existingRoleUser, err := s.roleUserRepository.GetByUserIdAndRoleId(ctx, userExisting.ID, userExisting.RoleID)
		if existingRoleUser == nil || err != nil {
			log.Errorf("user role data is not exist ID = %d RoleId = %d", userExisting.ID, userExisting.RoleID)
			return exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataUpdateFailed, err)
		}

		userExisting.Name = request.Name
		userExisting.Email = request.Email
		userExisting.Password = string(hashedPassword)
		userExisting.Avatar = ""

		// update user:
		userResult, err := s.userRepository.UpdateTx(ctx, tx.Client(), userExisting)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
		}

		existingRoleUser.UserID = userResult.ID
		existingRoleUser.RoleID = request.RoleId

		_, err = s.roleUserRepository.UpdateTx(ctx, tx.Client(), existingRoleUser)
		if err != nil {
			return exceptions.NewBusinessLogicError(exceptions.DataUpdateFailed, err)
		}

		// set value to userUpdated for return value:
		userUpdated = userResult

		// create cache, don't throw exception if failed:
		_, _ = s.cache.Create(ctx, CacheKeyUserWithId(id), userUpdated, vars.GetTtlShortPeriod())

		return nil

	}); err != nil {
		// add rollback logic here
		log.Error("do rollback from transactional database operation")
		return nil, err
	}

	return userUpdated, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id uint64) (*ent.User, error) {
	data, err := s.userRepository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataDeleteFailed, err)
	}

	_, _ = s.cache.Delete(ctx, CacheKeyUserWithId(id))

	return data, nil
}

func (s *UserServiceImpl) GetById(ctx context.Context, id uint64) (*ent.User, error) {

	userCache, _ := s.cache.Get(ctx, CacheKeyUserWithId(id), &ent.User{})
	if u, ok := userCache.(*ent.User); ok && u != nil {
		return u, nil
	}

	result, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	_, _ = s.cache.Create(ctx, CacheKeyUserWithId(id), result, vars.GetTtlShortPeriod())

	return result, nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]*ent.User, error) {
	userCache, _ := s.cache.Get(ctx, CacheKeyUsers(), &[]*ent.User{})
	if cached, ok := userCache.(*[]*ent.User); ok && cached != nil {
		userResult := append([]*ent.User(nil), *cached...)
		return userResult, nil
	}

	result, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	_, _ = s.cache.Create(ctx, CacheKeyUsers(), &result, vars.GetTtlShortPeriod())

	return result, nil
}
