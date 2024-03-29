package registry

import (
	"myapp/configs/rabbitmq/connection"
	"myapp/internal/component/rabbitmq/config"
)

type ProducerRegistry struct {
	conn *connection.RabbitMQConnection
}

func NewProducerRegistry(conn *connection.RabbitMQConnection) *ProducerRegistry {
	return &ProducerRegistry{conn: conn}
}

func (f *ProducerRegistry) Register() {

	mqConfigs := []config.IRabbitMQConfig{
		config.NewRabbitMQConfigExample(),

		//add config here ...
	}

	// run registry:
	f.execute(mqConfigs)
}
