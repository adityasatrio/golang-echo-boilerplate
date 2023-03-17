package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/role/dto"
	"myapp/internal/applications/role/repository"
	"myapp/internal/applications/transaction"
)

type RoleServiceImpl struct {
	repository  repository.RoleRepository
	transaction *transaction.TxService
}

func NewRoleServiceImpl(repository repository.RoleRepository, transaction *transaction.TxService) *RoleServiceImpl {
	return &RoleServiceImpl{repository: repository, transaction: transaction}
}

func (s *RoleServiceImpl) Create(ctx context.Context, request dto.RoleRequest) (*ent.Role, error) {
	role, err := s.repository.Create(ctx, request)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10003, err)
	}

	return role, nil
}

func (s *RoleServiceImpl) Update(ctx context.Context, request dto.RoleRequest, id uint64) (*ent.Role, error) {
	role, err := s.repository.Update(ctx, request, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10004, err)
	}

	return role, nil
}

func (s *RoleServiceImpl) SoftDelete(ctx context.Context, id uint64) (*ent.Role, error) {
	data, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10005, err)
	}

	return data, nil
}

func (s *RoleServiceImpl) Delete(ctx context.Context, id uint64) (*ent.Role, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	_, err = s.repository.Delete(ctx, exist.ID)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return exist, nil
}

func (s *RoleServiceImpl) GetById(ctx context.Context, id uint64) (*ent.Role, error) {
	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}

func (s *RoleServiceImpl) GetAll(ctx context.Context) ([]*ent.Role, error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.EBL10006, err)
	}

	return result, nil
}
