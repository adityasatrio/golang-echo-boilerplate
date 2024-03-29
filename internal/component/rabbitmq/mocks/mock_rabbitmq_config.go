package mocks

import "github.com/stretchr/testify/mock"

type MockRabbitMQConfig struct {
	mock.Mock
}

func (m *MockRabbitMQConfig) GetExchangeKind() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetExchangeDirect() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetQueueDirect() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetRoutingKeyDirect() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetExchangeDlx() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetQueueDlq() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetRoutingKeyDlx() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetExchangeJunk() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetQueueJunk() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetRoutingKeyJunk() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRabbitMQConfig) GetTtl() int64 {
	args := m.Called()
	return int64(args.Int(0))
}

func (m *MockRabbitMQConfig) GetDelay() int64 {
	args := m.Called()
	return int64(args.Int(0))
}

func (m *MockRabbitMQConfig) GetLimit() int64 {
	args := m.Called()
	return int64(args.Int(0))
}
