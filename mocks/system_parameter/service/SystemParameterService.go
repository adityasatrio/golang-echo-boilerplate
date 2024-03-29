// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_service

import (
	context "context"
	ent "myapp/ent"
	dto "myapp/internal/applications/system_parameter/dto"

	mock "github.com/stretchr/testify/mock"
)

// SystemParameterService is an autogenerated mock type for the SystemParameterService type
type SystemParameterService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, create
func (_m *SystemParameterService) Create(ctx context.Context, create *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error) {
	ret := _m.Called(ctx, create)

	var r0 *ent.SystemParameter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.SystemParameterCreateRequest) (*ent.SystemParameter, error)); ok {
		return rf(ctx, create)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.SystemParameterCreateRequest) *ent.SystemParameter); ok {
		r0 = rf(ctx, create)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.SystemParameter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.SystemParameterCreateRequest) error); ok {
		r1 = rf(ctx, create)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SystemParameterService) Delete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.SystemParameter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*ent.SystemParameter, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.SystemParameter); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.SystemParameter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *SystemParameterService) GetAll(ctx context.Context) ([]*ent.SystemParameter, error) {
	ret := _m.Called(ctx)

	var r0 []*ent.SystemParameter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*ent.SystemParameter, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*ent.SystemParameter); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.SystemParameter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *SystemParameterService) GetById(ctx context.Context, id int) (*ent.SystemParameter, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.SystemParameter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*ent.SystemParameter, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.SystemParameter); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.SystemParameter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SoftDelete provides a mock function with given fields: ctx, id
func (_m *SystemParameterService) SoftDelete(ctx context.Context, id int) (*ent.SystemParameter, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.SystemParameter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*ent.SystemParameter, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.SystemParameter); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.SystemParameter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, update
func (_m *SystemParameterService) Update(ctx context.Context, id int, update *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error) {
	ret := _m.Called(ctx, id, update)

	var r0 *ent.SystemParameter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *dto.SystemParameterUpdateRequest) (*ent.SystemParameter, error)); ok {
		return rf(ctx, id, update)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, *dto.SystemParameterUpdateRequest) *ent.SystemParameter); ok {
		r0 = rf(ctx, id, update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.SystemParameter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, *dto.SystemParameterUpdateRequest) error); ok {
		r1 = rf(ctx, id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSystemParameterService creates a new instance of SystemParameterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSystemParameterService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SystemParameterService {
	mock := &SystemParameterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
