package registry

import (
	"github.com/labstack/gommon/log"
	"myapp/configs/rabbitmq/connection"
	"myapp/ent"
	example "myapp/internal/applications/example/rabbitmq"
	"myapp/internal/component/rabbitmq/config"
)

type ConsumerRegistry struct {
	client *ent.Client
	conn   *connection.RabbitMQConnection
}

func NewConsumerRegistry(client *ent.Client, conn *connection.RabbitMQConnection) *ConsumerRegistry {
	return &ConsumerRegistry{client: client, conn: conn}
}

func (f *ConsumerRegistry) Register() {

	//init testing inbound:
	inbound := example.InitializedExampleInbound(f.client, f.conn)
	_, err := inbound.GetMessage(config.NewRabbitMQConfigExample())
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	//init other consumer here...

}
