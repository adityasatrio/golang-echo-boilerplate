package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"myapp/ent"
	"myapp/internal/applications/role/dto"
	mockRoleRepo "myapp/mocks/role/repository"
	mockTrx "myapp/mocks/transaction"
	"testing"
)

var mockRoleRepository = new(mockRoleRepo.RoleRepository)
var mockTransaction = new(mockTrx.TrxService)
var service = NewRoleServiceImpl(mockRoleRepository, mockTransaction)

func getRole(id uint64, name string, text string) ent.Role {
	return ent.Role{
		ID:   id,
		Name: name,
		Text: text,
	}
}

func TestRoleServiceImpl_Create(t *testing.T) {
	ctx := context.Background()
	id := uint64(0)
	roleRequest := dto.RoleRequest{
		Name: "CS",
		Text: "Customer Service",
	}

	t.Run("Create_success", func(t *testing.T) {
		role := getRole(id, "CS", "Customer Service")
		mockRoleRepository.On("Create", ctx, role).Return(&role, nil).Once()

		result, err := service.Create(ctx, roleRequest)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Create_failed", func(t *testing.T) {
		errMessage := errors.New("failed create role")
		role := getRole(id, "CS", "Customer Service")
		mockRoleRepository.On("Create", ctx, role).Return(nil, errMessage).Once()

		result, err := service.Create(ctx, roleRequest)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestRoleServiceImpl_Update(t *testing.T) {
	ctx := context.Background()
	id := uint64(0)

	roleRequest := dto.RoleRequest{
		Name: "CS",
		Text: "Customer Service",
	}

	t.Run("Update_success", func(t *testing.T) {
		role := getRole(id, "CS", "Customer Service")
		mockRoleRepository.On("Update", ctx, role, id).Return(&role, nil).Once()

		result, err := service.Update(ctx, roleRequest, id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Update_failed", func(t *testing.T) {
		errMessage := errors.New("failed update role")
		role := getRole(id, "CS", "Customer Service")
		mockRoleRepository.On("Update", ctx, role, id).Return(nil, errMessage).Once()

		result, err := service.Update(ctx, roleRequest, id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestRoleServiceImpl_SoftDelete(t *testing.T) {
	ctx := context.Background()
	id := uint64(100)

	t.Run("SoftDelete_success", func(t *testing.T) {
		role := getRole(id, "CS", "Customer Service")
		mockRoleRepository.On("SoftDelete", ctx, id).Return(&role, nil).Once()

		result, err := service.SoftDelete(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("SoftDelete_failed", func(t *testing.T) {
		errMessage := errors.New("failed soft delete role")
		mockRoleRepository.On("SoftDelete", ctx, id).Return(nil, errMessage).Once()

		result, err := service.SoftDelete(ctx, id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestRoleServiceImpl_Delete(t *testing.T) {
	ctx := context.Background()
	id := uint64(100)
	t.Run("Delete_success", func(t *testing.T) {

		role := getRole(id, "Doctor", "Expert")

		mockRoleRepository.On("GetById", ctx, id).Return(&role, nil).Once()
		mockRoleRepository.On("Delete", ctx, id).Return(&role, nil).Once()
		result, err := service.Delete(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Delete_failed_WhenDataNotFound", func(t *testing.T) {
		errMessage := errors.New("failed delete role")
		mockRoleRepository.On("GetById", ctx, id).Return(nil, errMessage).Once()

		result, err := service.Delete(ctx, id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("Delete_failed_WhenDeleteFailed", func(t *testing.T) {
		errMessage := errors.New("failed delete role")
		role := getRole(id, "Doctor", "Expert")
		mockRoleRepository.On("GetById", ctx, id).Return(&role, nil).Once()
		mockRoleRepository.On("Delete", ctx, id).Return(nil, errMessage).Once()

		result, err := service.Delete(ctx, id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestRoleServiceImpl_GetById(t *testing.T) {
	ctx := context.Background()
	id := uint64(100)

	t.Run("GetById_success", func(t *testing.T) {
		role := getRole(id, "CS", "Customer Service")
		mockRoleRepository.On("GetById", ctx, id).Return(&role, nil).Once()

		result, err := service.GetById(ctx, id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetById_failed", func(t *testing.T) {
		errMessage := errors.New("failed get id role")
		mockRoleRepository.On("GetById", ctx, id).Return(nil, errMessage).Once()

		result, err := service.GetById(ctx, id)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func TestRoleServiceImpl_GetAll(t *testing.T) {
	ctx := context.Background()
	t.Run("GetAll_success", func(t *testing.T) {
		role := getRole(uint64(1), "Doctor", "Expert")
		mockListRole := make([]*ent.Role, 0)
		mockListRole = append(mockListRole, &role)
		mockRoleRepository.On("GetAll", ctx).Return(mockListRole, nil).Once()

		result, err := service.GetAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("GetAll_failed", func(t *testing.T) {
		errMessage := errors.New("failed get all role")
		mockRoleRepository.On("GetAll", ctx).Return(nil, errMessage).Once()
		result, err := service.GetAll(ctx)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}
