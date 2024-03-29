package producer

import "myapp/internal/component/rabbitmq/config"

type Producer interface {
	SendToDirect(producer config.IRabbitMQConfig, message []byte) (bool, error)
	SendToJunk(producer config.IRabbitMQConfig, message []byte) (bool, error)
}
