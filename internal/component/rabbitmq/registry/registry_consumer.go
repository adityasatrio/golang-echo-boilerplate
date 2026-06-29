package registry

import (
	"github.com/labstack/gommon/log"
	"myapp/internal/builder"
	"myapp/internal/component/rabbitmq/config"
)

type ConsumerRegistry struct {
	container *builder.Container
}

func NewConsumerRegistry(container *builder.Container) *ConsumerRegistry {
	return &ConsumerRegistry{container: container}
}

func (f *ConsumerRegistry) Register() {

	// init testing inbound:
	inbound := f.container.BuildExampleRabbitMQInbound()
	_, err := inbound.GetMessage(config.NewRabbitMQConfigExample())
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// init other consumer here...

}
