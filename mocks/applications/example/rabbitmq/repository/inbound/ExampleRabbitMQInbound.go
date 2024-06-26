// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_inbound

import (
	config "myapp/internal/component/rabbitmq/config"

	mock "github.com/stretchr/testify/mock"
)

// ExampleRabbitMQInbound is an autogenerated mock type for the ExampleRabbitMQInbound type
type ExampleRabbitMQInbound struct {
	mock.Mock
}

// GetMessage provides a mock function with given fields: cfg
func (_m *ExampleRabbitMQInbound) GetMessage(cfg config.IRabbitMQConfig) (bool, error) {
	ret := _m.Called(cfg)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig) (bool, error)); ok {
		return rf(cfg)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig) bool); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig) error); ok {
		r1 = rf(cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExampleRabbitMQInbound creates a new instance of ExampleRabbitMQInbound. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExampleRabbitMQInbound(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExampleRabbitMQInbound {
	mock := &ExampleRabbitMQInbound{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
