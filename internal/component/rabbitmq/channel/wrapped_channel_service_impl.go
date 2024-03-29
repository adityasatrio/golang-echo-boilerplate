package channel

import (
	"context"
	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"myapp/configs/rabbitmq/connection"
)

type WrappedChannelServiceImpl struct {
	connection *connection.RabbitMQConnection
}

func NewWrappedChannel(connection *connection.RabbitMQConnection) *WrappedChannelServiceImpl {
	return &WrappedChannelServiceImpl{connection: connection}
}

func (wc *WrappedChannelServiceImpl) PublishMessage(exchange, key string, msg amqp.Publishing) error {
	log.Debug("Performing additional actions before produce...")
	return wc.connection.GetChannel().PublishWithContext(context.Background(), exchange, key, false, false, msg)
}

func (wc *WrappedChannelServiceImpl) ConsumeMessage(queue string) (<-chan amqp.Delivery, error) {
	log.Debug("Performing additional actions before consume...")
	return wc.connection.GetChannel().Consume(queue, "", false, false, false, false, nil)
}
