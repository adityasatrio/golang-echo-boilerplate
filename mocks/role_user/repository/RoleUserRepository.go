// Code generated by mockery v2.23.0. DO NOT EDIT.

package mock_repository

import (
	context "context"
	ent "myapp/ent"

	mock "github.com/stretchr/testify/mock"
)

// RoleUserRepository is an autogenerated mock type for the RoleUserRepository type
type RoleUserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, client, request
func (_m *RoleUserRepository) Create(ctx context.Context, client *ent.Tx, request ent.RoleUser) (*ent.RoleUser, error) {
	ret := _m.Called(ctx, client, request)

	var r0 *ent.RoleUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, ent.RoleUser) (*ent.RoleUser, error)); ok {
		return rf(ctx, client, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, ent.RoleUser) *ent.RoleUser); ok {
		r0 = rf(ctx, client, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.RoleUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Tx, ent.RoleUser) error); ok {
		r1 = rf(ctx, client, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByUserId provides a mock function with given fields: ctx, client, id
func (_m *RoleUserRepository) DeleteByUserId(ctx context.Context, client *ent.Tx, id uint64) (int, error) {
	ret := _m.Called(ctx, client, id)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, uint64) (int, error)); ok {
		return rf(ctx, client, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, uint64) int); ok {
		r0 = rf(ctx, client, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Tx, uint64) error); ok {
		r1 = rf(ctx, client, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, client, request, id
func (_m *RoleUserRepository) Update(ctx context.Context, client *ent.Tx, request ent.RoleUser, id uint64) (*ent.RoleUser, error) {
	ret := _m.Called(ctx, client, request, id)

	var r0 *ent.RoleUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, ent.RoleUser, uint64) (*ent.RoleUser, error)); ok {
		return rf(ctx, client, request, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, ent.RoleUser, uint64) *ent.RoleUser); ok {
		r0 = rf(ctx, client, request, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.RoleUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Tx, ent.RoleUser, uint64) error); ok {
		r1 = rf(ctx, client, request, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRoleUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleUserRepository creates a new instance of RoleUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleUserRepository(t mockConstructorTestingTNewRoleUserRepository) *RoleUserRepository {
	mock := &RoleUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}