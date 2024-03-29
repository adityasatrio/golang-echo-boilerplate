package producer

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"myapp/internal/component/rabbitmq/mocks"
	"myapp/internal/component/rabbitmq/utils"
	mock_channel "myapp/mocks/component/rabbitmq/channel"
	"testing"
)

func TestSendToDirect_Success(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockExchange := "exchange_direct"
	mockRoutingKey := "routing_key_direct"
	mockMessage := []byte("Test message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
	}

	mockConfig.On("GetExchangeDirect").Return(mockExchange)
	mockConfig.On("GetRoutingKeyDirect").Return(mockRoutingKey)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).Return(nil)

	// Call the function being tested
	success, err := producerService.SendToDirect(mockConfig, mockMessage)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, success)
	mockChannel.AssertExpectations(t)
}

func TestSendToDirect_Failed(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockExchange := "exchange_direct"
	mockRoutingKey := "routing_key_direct"
	mockMessage := []byte("Test message")
	mockError := errors.New("error sending message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
	}

	mockConfig.On("GetExchangeDirect").Return(mockExchange)
	mockConfig.On("GetRoutingKeyDirect").Return(mockRoutingKey)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).
		Return(mockError)

	// Call the function being tested
	success, err := producerService.SendToDirect(mockConfig, mockMessage)

	// Assertions
	assert.Error(t, err)
	assert.False(t, success)
	assert.EqualError(t, err, "error sending message")
	mockChannel.AssertExpectations(t)
}

func TestSendToJunk_Success(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockExchange := "exchange_junk"
	mockRoutingKey := "routing_key_junk"
	mockMessage := []byte("Test message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
	}

	mockConfig.On("GetExchangeJunk").Return(mockExchange)
	mockConfig.On("GetRoutingKeyJunk").Return(mockRoutingKey)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).Return(nil)

	// Call the function being tested
	success, err := producerService.SendToJunk(mockConfig, mockMessage)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, success)
	mockChannel.AssertExpectations(t)
}

func TestSendToJunk_Failed(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockExchange := "exchange_junk"
	mockRoutingKey := "routing_key_junk"
	mockMessage := []byte("Test message")
	mockError := errors.New("error sending message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
	}

	mockConfig.On("GetExchangeJunk").Return(mockExchange)
	mockConfig.On("GetRoutingKeyJunk").Return(mockRoutingKey)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).
		Return(mockError)

	// Call the function being tested
	success, err := producerService.SendToJunk(mockConfig, mockMessage)

	// Assertions
	assert.Error(t, err)
	assert.False(t, success)
	assert.EqualError(t, err, "error sending message")
	mockChannel.AssertExpectations(t)
}
