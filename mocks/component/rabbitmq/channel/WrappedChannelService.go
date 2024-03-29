// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_channel

import (
	amqp091 "github.com/rabbitmq/amqp091-go"

	mock "github.com/stretchr/testify/mock"
)

// WrappedChannelService is an autogenerated mock type for the WrappedChannelService type
type WrappedChannelService struct {
	mock.Mock
}

// ConsumeMessage provides a mock function with given fields: queue
func (_m *WrappedChannelService) ConsumeMessage(queue string) (<-chan amqp091.Delivery, error) {
	ret := _m.Called(queue)

	var r0 <-chan amqp091.Delivery
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (<-chan amqp091.Delivery, error)); ok {
		return rf(queue)
	}
	if rf, ok := ret.Get(0).(func(string) <-chan amqp091.Delivery); ok {
		r0 = rf(queue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp091.Delivery)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(queue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishMessage provides a mock function with given fields: exchange, key, msg
func (_m *WrappedChannelService) PublishMessage(exchange string, key string, msg amqp091.Publishing) error {
	ret := _m.Called(exchange, key, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, amqp091.Publishing) error); ok {
		r0 = rf(exchange, key, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWrappedChannelService creates a new instance of WrappedChannelService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWrappedChannelService(t interface {
	mock.TestingT
	Cleanup(func())
}) *WrappedChannelService {
	mock := &WrappedChannelService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
