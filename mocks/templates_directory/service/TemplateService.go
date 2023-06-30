// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_service

import (
	context "context"
	ent "myapp/ent"
	dto "myapp/internal/applications/A_templates_directory/dto"

	mock "github.com/stretchr/testify/mock"
)

// TemplateService is an autogenerated mock type for the TemplateService type
type TemplateService struct {
	mock.Mock
}

// LogicFunction provides a mock function with given fields: ctx, request
func (_m *TemplateService) LogicFunction(ctx context.Context, request *dto.ExampleRequest) (*ent.Pet, error) {
	ret := _m.Called(ctx, request)

	var r0 *ent.Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ExampleRequest) (*ent.Pet, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ExampleRequest) *ent.Pet); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Pet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ExampleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTemplateService creates a new instance of TemplateService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTemplateService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TemplateService {
	mock := &TemplateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}