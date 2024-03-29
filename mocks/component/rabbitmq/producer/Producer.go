// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_producer

import (
	config "myapp/internal/component/rabbitmq/config"

	mock "github.com/stretchr/testify/mock"
)

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// SendToDirect provides a mock function with given fields: _a0, message
func (_m *Producer) SendToDirect(_a0 config.IRabbitMQConfig, message []byte) (bool, error) {
	ret := _m.Called(_a0, message)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) (bool, error)); ok {
		return rf(_a0, message)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) bool); ok {
		r0 = rf(_a0, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig, []byte) error); ok {
		r1 = rf(_a0, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendToJunk provides a mock function with given fields: _a0, message
func (_m *Producer) SendToJunk(_a0 config.IRabbitMQConfig, message []byte) (bool, error) {
	ret := _m.Called(_a0, message)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) (bool, error)); ok {
		return rf(_a0, message)
	}
	if rf, ok := ret.Get(0).(func(config.IRabbitMQConfig, []byte) bool); ok {
		r0 = rf(_a0, message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(config.IRabbitMQConfig, []byte) error); ok {
		r1 = rf(_a0, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProducer creates a new instance of Producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProducer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Producer {
	mock := &Producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
