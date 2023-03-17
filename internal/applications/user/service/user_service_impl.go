package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/transaction"
	"myapp/internal/applications/user/dto"
	"myapp/internal/applications/user/repository"
)

type UserServiceImpl struct {
	repository  repository.UserRepository
	transaction *transaction.TxService
}

func NewUserServiceImpl(repository repository.UserRepository, transaction *transaction.TxService) *UserServiceImpl {
	return &UserServiceImpl{repository: repository, transaction: transaction}
}

func (s *UserServiceImpl) Create(ctx context.Context, request dto.UserRequest) (*ent.User, error) {
	data, err := s.repository.Create(ctx, request)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
	}

	return data, nil
}

func (s *UserServiceImpl) Update(ctx context.Context, id uint64, request dto.UserRequest) (*ent.User, error) {
	role, err := s.repository.Update(ctx, request, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10004, err)
	}

	return role, nil
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
