package inbound

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/component/rabbitmq/mocks"
	mock_service "myapp/mocks/applications/example/rabbitmq/service"
	mock_channel "myapp/mocks/component/rabbitmq/channel"
	mock_producer "myapp/mocks/component/rabbitmq/producer"
	"testing"
	"time"
)

func TestParsingData(t *testing.T) {
	msg := []byte(`{"key": "example_rabbit_key", "value":"example_rabbit_values"}`)
	data := parsingData(msg)
	assert.Equal(t, "example_rabbit_key", data.Key)
	assert.Equal(t, "example_rabbit_values", data.Value)
}

func TestParsingData_Failed(t *testing.T) {
	msg := []byte(nil)
	data := parsingData(msg)
	assert.Equal(t, "", data.Key)
	assert.Equal(t, "", data.Value)
}

func TestExampleInbound_GetMessage_Failed_GetConsumer(t *testing.T) {
	mockConfig := new(mocks.MockRabbitMQConfig)
	mockChannel := new(mock_channel.WrappedChannelService)
	mockExampleService := new(mock_service.ExampleRabbitMQService)
	mockProducer := new(mock_producer.Producer)
	exampleInbound := NewExampleRabbitMQInbound(mockChannel, mockExampleService, mockProducer)

	// Prepare test data
	//example request:
	request := dto.SystemParameterCreateRequest{
		Key:   "example_rabbit_key",
		Value: "example_rabbit_values",
	}

	//parsing to json:
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		jsonMessage = []byte("")
	}
	var messageBody = mocks.MockDeliveryChannel(jsonMessage)

	mockConfig.On("GetQueueDirect").Return("direct_queue")

	mockChannel.On("ConsumeMessage", "direct_queue").
		Return(messageBody, errors.New("failed execute consumer"))

	result, err := exampleInbound.GetMessage(mockConfig)
	assert.NotNil(t, err)
	assert.Equal(t, false, result)
}

func TestExampleInbound_GetMessage(t *testing.T) {
	mockConfig := new(mocks.MockRabbitMQConfig)
	mockChannel := new(mock_channel.WrappedChannelService)
	mockExampleService := new(mock_service.ExampleRabbitMQService)
	mockProducer := new(mock_producer.Producer)
	exampleInbound := NewExampleRabbitMQInbound(mockChannel, mockExampleService, mockProducer)

	// Prepare test data
	//example request:
	request := dto.SystemParameterCreateRequest{
		Key:   "example_rabbit_key",
		Value: "example_rabbit_values",
	}

	//parsing to json:
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		jsonMessage = []byte("")
	}
	var messageBody = mocks.MockDeliveryChannel(jsonMessage)

	mockConfig.On("GetQueueDirect").Return("direct_queue")

	mockChannel.On("ConsumeMessage", "direct_queue").
		Return(messageBody, nil)

	mockExampleService.On("GetMessage", context.Background(), &request).
		Return(nil, nil)

	// Use a channel to signal when the goroutine completes
	done := make(chan struct{})
	go func(msg <-chan amqp.Delivery) {
		defer close(done)

		// Lakukan pemrosesan pesan
		_, err := exampleInbound.GetMessage(mockConfig)
		if err != nil {
			return
		}
		time.Sleep(5 * time.Second)
	}(messageBody)

	// Wait for the goroutines to complete or timeout after a certain duration
	select {
	case <-done:
	case <-time.After(60 * time.Second): // Adjust the duration as needed
		t.Error("Test timed out")
	}

	mockChannel.AssertExpectations(t)
	mockExampleService.AssertExpectations(t)
}

func TestExampleInbound_GetMessage_Failed(t *testing.T) {
	mockConfig := new(mocks.MockRabbitMQConfig)
	mockChannel := new(mock_channel.WrappedChannelService)

	mockExampleService := new(mock_service.ExampleRabbitMQService)
	mockProducer := new(mock_producer.Producer)
	exampleInbound := NewExampleRabbitMQInbound(mockChannel, mockExampleService, mockProducer)

	// Prepare test data
	//example request:
	request := dto.SystemParameterCreateRequest{
		Key:   "example_rabbit_key",
		Value: "example_rabbit_values",
	}

	//parsing to json:
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		jsonMessage = []byte("")
	}
	var messageBody = mocks.MockDeliveryChannel(jsonMessage)

	mockConfig.On("GetQueueDirect").Return("direct_queue")
	mockConfig.On("GetLimit").Return(0)

	mockChannel.On("ConsumeMessage", "direct_queue").
		Return(messageBody, nil)

	mockExampleService.On("GetMessage", context.Background(), &request).
		Return(nil, errors.New("failed execute service"))

	mockProducer.On("SendToJunk", mockConfig, jsonMessage).
		Return(true, nil)

	// Use a channel to signal when the goroutine completes
	done := make(chan struct{})
	go func(msg <-chan amqp.Delivery) {
		defer close(done)

		// Lakukan pemrosesan pesan
		_, err := exampleInbound.GetMessage(mockConfig)
		if err != nil {
			return
		}
		time.Sleep(5 * time.Second)
	}(messageBody)

	// Wait for the goroutines to complete or timeout after a certain duration
	select {
	case <-done:
	case <-time.After(60 * time.Second): // Adjust the duration as needed
		t.Error("Test timed out")
	}

	mockChannel.AssertExpectations(t)
	mockExampleService.AssertExpectations(t)
	mockProducer.AssertExpectations(t)
}

func TestExampleInbound_GetMessage_Failed_Send_Junk(t *testing.T) {
	mockConfig := new(mocks.MockRabbitMQConfig)
	mockChannel := new(mock_channel.WrappedChannelService)
	mockExampleService := new(mock_service.ExampleRabbitMQService)
	mockProducer := new(mock_producer.Producer)
	exampleInbound := NewExampleRabbitMQInbound(mockChannel, mockExampleService, mockProducer)

	// Prepare test data
	//example request:
	request := dto.SystemParameterCreateRequest{
		Key:   "example_rabbit_key",
		Value: "example_rabbit_values",
	}

	//parsing to json:
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		jsonMessage = []byte("")
	}
	var messageBody = mocks.MockDeliveryChannel(jsonMessage)

	mockConfig.On("GetQueueDirect").Return("direct_queue")
	mockConfig.On("GetLimit").Return(0)

	mockChannel.On("ConsumeMessage", "direct_queue").
		Return(messageBody, nil)

	mockExampleService.On("GetMessage", context.Background(), &request).
		Return(nil, errors.New("failed execute service"))

	mockProducer.On("SendToJunk", mockConfig, jsonMessage).
		Return(true, errors.New("failed execute send junk"))

	// Use a channel to signal when the goroutine completes
	done := make(chan struct{})
	go func(msg <-chan amqp.Delivery) {
		defer close(done)

		// Lakukan pemrosesan pesan
		_, err := exampleInbound.GetMessage(mockConfig)
		if err != nil {
			return
		}
		time.Sleep(5 * time.Second)
	}(messageBody)

	// Wait for the goroutines to complete or timeout after a certain duration
	select {
	case <-done:
	case <-time.After(60 * time.Second): // Adjust the duration as needed
		t.Error("Test timed out")
	}

	mockChannel.AssertExpectations(t)
	mockExampleService.AssertExpectations(t)
	mockProducer.AssertExpectations(t)
}
