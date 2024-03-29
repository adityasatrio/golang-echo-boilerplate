package channel

import amqp "github.com/rabbitmq/amqp091-go"

type WrappedChannelService interface {
	PublishMessage(exchange, key string, msg amqp.Publishing) error
	ConsumeMessage(queue string) (<-chan amqp.Delivery, error)
}
