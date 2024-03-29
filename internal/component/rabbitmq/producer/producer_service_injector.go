//go:build wireinject
// +build wireinject

package producer

import (
	"github.com/google/wire"
	"myapp/configs/rabbitmq/connection"
	"myapp/internal/component/rabbitmq/channel"
)

var provider = wire.NewSet(
	NewProducerService,
	channel.NewWrappedChannel,

	wire.Bind(new(channel.WrappedChannelService), new(*channel.WrappedChannelServiceImpl)),
	wire.Bind(new(Producer), new(*ProducerServiceImpl)),
)

func InitializedProducer(connection *connection.RabbitMQConnection) *ProducerServiceImpl {
	wire.Build(provider)
	return nil
}
