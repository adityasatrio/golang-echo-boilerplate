package inbound

import "myapp/internal/component/rabbitmq/config"

type ExampleRabbitMQInbound interface {
	GetMessage(cfg config.IRabbitMQConfig) (bool, error)
}
