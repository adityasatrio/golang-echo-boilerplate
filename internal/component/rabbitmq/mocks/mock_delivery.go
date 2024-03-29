package mocks

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func MockDeliveryChannel(message []byte) <-chan amqp.Delivery {
	deliveryChannel := make(chan amqp.Delivery)

	go func() {
		// Simulating AMQP message deliveries
		for i := 1; i <= 1; i++ {
			message := amqp.Delivery{
				Body:         message,
				DeliveryMode: amqp.Transient,
				DeliveryTag:  uint64(i),
				Acknowledger: &MockAcknowledger{},
			}
			// Simulate delivering a message to the channel
			deliveryChannel <- message
			time.Sleep(5 * time.Second) // Simulate delay between messages
		}

		// Close the channel after delivering all messages
		close(deliveryChannel)
	}()

	return deliveryChannel
}
