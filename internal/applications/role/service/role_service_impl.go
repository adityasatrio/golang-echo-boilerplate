package service

import (
	"context"
	"myapp/ent"
	"myapp/exceptions"
	"myapp/internal/applications/role/dto"
	"myapp/internal/applications/role/repository"
)

type RoleServiceImpl struct {
	repository repository.RoleRepository
}

func NewRoleService(repository repository.RoleRepository) *RoleServiceImpl {
	return &RoleServiceImpl{repository: repository}
}

func (s *RoleServiceImpl) Create(ctx context.Context, request dto.RoleRequest) (*ent.Role, error) {

	roleReq := ent.Role{
		Name: request.Name,
		Text: request.Text,
	}

	roleSaved, err := s.repository.Create(ctx, roleReq)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataCreateFailed, err)
	}

	return roleSaved, nil
}

func (s *RoleServiceImpl) Update(ctx context.Context, request dto.RoleRequest, id uint64) (*ent.Role, error) {

	existing, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	existing.Name = request.Name
	existing.Text = request.Text

	roleUpdated, err := s.repository.Update(ctx, existing)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataUpdateFailed, err)
	}

	return roleUpdated, nil
}

func (s *RoleServiceImpl) SoftDelete(ctx context.Context, id uint64) (*ent.Role, error) {
	data, err := s.repository.SoftDelete(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataDeleteFailed, err)
	}

	return data, nil
}

func (s *RoleServiceImpl) Delete(ctx context.Context, id uint64) (*ent.Role, error) {
	exist, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataNotFound, err)
	}

	_, err = s.repository.Delete(ctx, exist.ID)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	return exist, nil
}

func (s *RoleServiceImpl) GetById(ctx context.Context, id uint64) (*ent.Role, error) {
	result, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	return result, nil
}

func (s *RoleServiceImpl) GetAll(ctx context.Context) ([]*ent.Role, error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, exceptions.NewBusinessLogicError(exceptions.DataGetFailed, err)
	}

	return result, nil
}
