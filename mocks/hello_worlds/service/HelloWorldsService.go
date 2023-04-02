// Code generated by mockery v2.23.0. DO NOT EDIT.

package mock_service

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HelloWorldsService is an autogenerated mock type for the HelloWorldsService type
type HelloWorldsService struct {
	mock.Mock
}

// Hello provides a mock function with given fields: ctx, message, errorFlag
func (_m *HelloWorldsService) Hello(ctx context.Context, message string, errorFlag string) (string, error) {
	ret := _m.Called(ctx, message, errorFlag)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, message, errorFlag)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, message, errorFlag)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, message, errorFlag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHelloWorldsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelloWorldsService creates a new instance of HelloWorldsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelloWorldsService(t mockConstructorTestingTNewHelloWorldsService) *HelloWorldsService {
	mock := &HelloWorldsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}